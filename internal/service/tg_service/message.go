package tg_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"myapp/internal/models"
	"myapp/pkg/files"
	my_regex "myapp/pkg/regex"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (srv *TgService) HandleMessage(m models.Update) error {
	msgText := m.Message.Text
	fromUsername := m.Message.From.UserName
	fromId := m.Message.From.Id
	srv.l.Info(fmt.Sprintf("HandleMessage: fromId-%d fromUsername-%s, msgText-%s", fromId, fromUsername, msgText))

	srv.SendMsgToServer(fromId, "user", msgText)

	if msgText == "/admin" {
		err := srv.M_admin(m)
		if err != nil {
			srv.SendMessage(fromId, ERR_MSG)
			srv.SendMessage(fromId, err.Error())
		}
		return err
	}

	user, err := srv.Db.GetUserById(fromId)
	if err != nil {
		return fmt.Errorf("HandleMessage GetUserById err: %v", err)
	}
	// if user.Id != 0 && user.Lives == 0 {
	// 	return nil
	// }

	if msgText == "/help" {
		srv.SendMessageAndDb(fromId, "@millioner_support\nвот контакт для связи")
		srv.Db.UpdateLatsActiontime(fromId)
		srv.Db.UpdateFeedbackTime(fromId)
		return nil
	}

	if user.IsLastPush == 1 {
		srv.SendMessageAndDb(fromId, "бот вам больше не доступен")
		return nil
	}

	if strings.HasPrefix(msgText, "/start") { // https://t.me/tgbotusername?start=ref01 -> /start ref01
		err := srv.M_start(m)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		srv.Db.UpdateFeedbackTime(fromId)
		return err
	}

	if strings.HasPrefix(msgText, "add_am_") { // add_am_1.1_
		animMessId := my_regex.GetStringInBetween(msgText, "add_am_", "_")
		if animMessId == "" {
			return fmt.Errorf("некоректный animMessId")
		}
		srv.Db.EditBotState(fromId, msgText)
		srv.SendMessage(fromId, fmt.Sprintf("Ожидание поста для animMessId %v", animMessId))
		return nil
	}

	err = srv.M_state(m)
	if err != nil {
		srv.SendMessageAndDb(fromId, ERR_MSG)
		srv.SendMessageAndDb(fromId, err.Error())
	}
	srv.Db.UpdateLatsActiontime(fromId)
	srv.Db.UpdateFeedbackTime(fromId)
	return err
}

func (srv *TgService) M_start(m models.Update) error {
	fromId := m.Message.Chat.Id
	msgText := m.Message.Text
	fromFirstName := m.Message.From.FirstName
	fromUsername := m.Message.From.UserName
	srv.l.Info(fmt.Sprintf("M_start: fromId: %d, fromUsername: %s, msgText: %s", fromId, fromUsername, msgText))

	refArr := strings.Split(msgText, " ")
	ref := ""
	if len(refArr) > 1 {
		ref = refArr[1]
	}

	// user, err := srv.Db.GetUserById(fromId)
	// if err != nil {
	// 	return fmt.Errorf("M_start GetUserById err: %v", err)
	// }
	// if user.CreatedAt != "" && srv.IsIgnoreUser(fromId) {
	// 	text := "К сожалению, время истекло и бот для вас больше недоступен.\nВы можете обратиться в поддержку через команду /help"
	// 	srv.SendMessageAndDb(fromId, text)
	// 	return nil
	// }

	err := srv.Db.AddNewUser(fromId, fromUsername, fromFirstName)
	if err != nil {
		return fmt.Errorf("M_start AddNewUser err: %v", err)
	}
	srv.Db.EditRef(fromId, ref)
	if fromId == 1394096901 {
		srv.Db.EditAdmin(fromId, 1)
	}
	srv.Db.EditBotState(fromId, "")
	srv.Db.EditLives(fromId, 3)
	srv.Db.EditStep(fromId, "1")
	srv.SendMessageAndDb(fromId, fmt.Sprintf("Привет, %s 👋", fromFirstName))
	// srv.SendAnimMessageHTML("1", fromId, animTimeout3000)

	// time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))

	// text := "Прямо сейчас начинай игру и забирай бонус 1000₽ за уверенный старт! 🚀"
	// replyMarkup := `{"inline_keyboard" : [
	// 	[{ "text": "Начать игру", "callback_data": "start_game" }]
	// ]}`
	// srv.SendMessageWRM(fromId, text, replyMarkup)
	
	// srv.SendMsgToServer(fromId, "bot", text)

	futureJson := map[string]string{
		"video_note":   fmt.Sprintf("@%s", "./files/krug_1.mp4"),
		"chat_id": strconv.Itoa(fromId),
	}
	cf, body, err := files.CreateForm(futureJson)
	if err != nil {
		return fmt.Errorf("HandleVideoNote CreateFormV2 err: %v", err)
	}
	srv.SendVideoNote(body, cf)

	srv.Db.EditBotState(fromId, "read_article_after_KNB_win")
	time.Sleep(time.Second*20)
	srv.SendMessage(fromId, "Введи кодовое слово ниже 👇🏻")

	return nil
}

func (srv *TgService) M_state(m models.Update) error {
	fromId := m.Message.Chat.Id
	msgText := m.Message.Text
	fromUsername := m.Message.From.UserName
	fromFirstName := m.Message.From.FirstName
	srv.l.Info(fmt.Sprintf("M_state: fromId: %d, fromUsername: %s, msgText: %s", fromId, fromUsername, msgText))

	user, err := srv.Db.GetUserById(fromId)
	if err != nil {
		srv.l.Warn(fmt.Errorf("M_state GetUserById err: %v", err))
	}
	srv.Db.UpdateLatsActiontime(fromId)
	if user.BotState == "" {
		return nil
	}

	if strings.HasPrefix(user.BotState, "add_am_") {
		animMessId := my_regex.GetStringInBetween(user.BotState, "add_am_", "_")
		animMess, err := srv.Db.GetAminMessByTxtId(animMessId)
		if err != nil {
			return fmt.Errorf("M_state GetAminMessByTxtId err: %v", err)
		}

		srv.l.Info("m.Message.Entities:", m.Message.Entities)
		srv.SendMessage(fromId, fmt.Sprintf("%+v", m.Message.Entities))

		// for _, v := range m.Message.Entities {
		// 	entityType := v.Type
		// 	entityStart := v.Offset
		// 	entityEnd := v.Offset + v.Length
		// 	var entityStartSymb string
		// 	var entityEndSymb string
		// 	if entityType == "bold" {
		// 		entityStartSymb = "<b>"
		// 		entityEndSymb = "</b>"
		// 	}
		// 	if entityType == "underline" {
		// 		entityStartSymb = "<u>"
		// 		entityEndSymb = "</u>"
		// 	}
		// 	for i := len([]rune(msgText)); i > 0; i-- {
		// 		if i == entityEnd 
		// 	}
		// }

		if animMess.TxtId != "" {
			err = srv.Db.EditAnimMessText(animMessId, msgText)
			if err != nil {
				return fmt.Errorf("M_state EditAnimMessText err: %v", err)
			}
			srv.SendMessage(fromId, "пост обновлен")
			srv.Db.EditBotState(fromId, "")
			return nil
		}
		err = srv.Db.AddNewAminMess(animMessId, msgText)
		if err != nil {
			return fmt.Errorf("M_state AddNewAminMess err: %v", err)
		}
		srv.SendMessage(fromId, "пост добавлен")
		srv.Db.EditBotState(fromId, "")
		return nil
	}

	if user.BotState == "read_article_after_KNB_win" { // Го, ко, коу, гоу, гэу
		if !strings.HasPrefix(strings.ToLower(msgText), "го") && !strings.HasPrefix(strings.ToLower(msgText), "ко") && !strings.HasPrefix(strings.ToLower(msgText), "гэ") && !strings.HasPrefix(strings.ToLower(msgText), "go") {
			srv.SendMessageAndDb(fromId, "❌ Вы неверно ввели кодовое слово, сверьтесь с лонгридом и попробуйте еще раз")
			return nil
		}

		srv.Db.EditBotState(fromId, "")
		// srv.SendAnimMessage("-1", fromId, animTimeout250)
		// srv.SendBalance(fromId, "30.000", animTimeout250)
		// srv.Db.EditStep(fromId, "9")
		// srv.SendAnimMessageHTML("9", fromId, animTimeoutTest)

		text := "Ну что, поехали, ответь правильно на 3 вопроса и уже сегодня сможешь заработать от 500.000₽ 😏"
		srv.SendMessage(fromId, text)
		err = srv.ShowMilQ(fromId, 1)
		if err != nil {
			return fmt.Errorf("M_state ShowMilQ err: %v", err)
		}

		// text := "Предлагаю тебе ответить на один вопрос 😏\nЗа него ты получишь +25.000₽ к банку💸"
		// replyMarkup :=`{"inline_keyboard" : [
		// 	[ { "text": "Давай попробуем", "callback_data": "show_q_3_" } ]
		// ]}`
		// srv.SendMessageWRM(fromId, text, replyMarkup)

		// srv.ShowMilQ(fromId, 2)
		// srv.Db.EditStep(fromId, "7")
		// srv.SendMsgToServer(fromId, "bot", text)
		return nil
	}

	if user.BotState == "wait_email" {
		msgTextEmail := msgText
		url := fmt.Sprintf("%s/api/v1/user?email=%s", srv.Cfg.ServerUrl, msgTextEmail)
		srv.l.Info("M_state wait_email иду к API", url)
		response, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("M_state wait_email Post err: %v", err)
		}
		srv.l.Info("M_state wait_email сходил к API")
		defer response.Body.Close()
	
		if response.StatusCode != http.StatusOK {
			bodyBytes, err := io.ReadAll(response.Body)
			if err != nil {
				return fmt.Errorf("M_state wait_email ReadAll err: %v", err)
			}
			return fmt.Errorf("M_state wait_email post %s bad response: [%d] %v", url, response.StatusCode, string(bodyBytes))
		}
	
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("M_state wait_email ReadAll err: %v", err)
		}
	
		resp := struct{
			Status string `json:"status"`
			Data   string `json:"data"`
		}{}
		json.Unmarshal(bodyBytes, &resp)
	
		if resp.Status == "success" {

			srv.Db.EditBotState(fromId, "")
			srv.Db.EditEmail(fromId, msgTextEmail)
			lichka, tgId,  _ := srv.GetLichka()
			srv.Db.EditLichka(fromId, lichka)
			// mess := fmt.Sprintf("Ваша личка %s", srv.AddAt(lichka))
			// srv.SendMessage(fromId, mess)

			url := fmt.Sprintf("%s/api/v1/lichka", srv.Cfg.ServerUrl)
			jsonBody := []byte(fmt.Sprintf(`{"lichka":"%s", "tg_id":"%d", "tg_username":"%s", "tg_name":"%s", "email":"%s"}`, lichka, tgId, fromUsername, fromFirstName, msgTextEmail))
			bodyReader := bytes.NewReader(jsonBody)
			_, err := http.Post(url, "application/json", bodyReader)
			if err != nil {
				return fmt.Errorf("M_state api/v1/lichka Post err: %v", err)
			}
			url = fmt.Sprintf("%s/api/v1/link_ref", srv.Cfg.ServerUrl)
			ref_id := srv.Refki[user.Ref]
			if ref_id != "хуй" {
				ref_id = "1000153272"
			}
			jsonBody = []byte(fmt.Sprintf(`{"user_email":"%s", "ref_id":"%s"}`, msgTextEmail, ref_id))
			bodyReader = bytes.NewReader(jsonBody)
			_, err = http.Post(url, "application/json", bodyReader)
			if err != nil {
				return fmt.Errorf("M_state api/v1/link_ref Post err: %v", err)
			}

			gifResp, _ := srv.CopyMessage(fromId, -1002074025173, 86) // https://t.me/c/2074025173/86
			// gifResp, _ := srv.SendVideoWCaption(fromId, "", "./files/gif_1.MOV")
			time.Sleep(time.Second*6)
			srv.DeleteMessage(fromId, gifResp.Result.MessageId)

			mess := "Молодчина! Тебе осталось выполнить последнее условие и ты наконец-то заберешь свою награду, благодаря которой заработаешь от 500.000₽ 🤑\n\nТебе нужно всего лишь прочитать текста, которые я для тебя подготовил и ответить правильно на вопросы после них😉\nДерзай 👇🏻"
			srv.SendMessage(fromId, mess)

			srv.Db.EditStep(fromId, "12")
			// srv.SendAnimMessageHTML("12", fromId, 2000)
			// text := "+45.000₽ уходят в твой банк за правильный ответ!💸\n\n🔐Чтобы разблокировать и забрать награду пришли мне кодовое слово из видео ☝🏻\n\n*Просмотр не займет много времени\nПосле пиши кодовое слово сюда.\nБуду ждать 👇🏻"
			// srv.SendVideoWCaption(fromId, text, "./files/VID_cod_1.mp4")
			// srv.CopyMessage(fromId, -1002074025173, 32)
			
			srv.SendAnimArticleHTMLV3("1.1", fromId, 2000)
			srv.CopyMessage(fromId, -1001998413789, 4) // https://t.me/c/1998413789/4
			srv.SendAnimArticleHTMLV3("1.2", fromId, 2000)
			srv.CopyMessage(fromId, -1001998413789, 9)

			text := `Какова главная причина того, что люди выходят из игры до достижения серьёзных результатов? Всё же настолько просто. 
		
			A) Недостаток мотивации и усердия.
			B) Отсутствие интереса к игре.
			C) Слишком сложные вопросы.
			D) Не хватает времени, чтобы играть.`
			reply_markup := `{"inline_keyboard" : [
				[ { "text": "A", "callback_data": "prodolzit_1_" }, { "text": "B", "callback_data": "__" }, { "text": "C", "callback_data": "__" }, { "text": "D", "callback_data": "__" }]
			]}`
			srv.SendMessageWRM(fromId, text, reply_markup)

			// srv.SendAnimMessageHTML("12", fromId, 2000)
			// // srv.Db.EditBotState(fromId, "read_article_after_TrurOrFalse_win")
			// srv.Db.EditBotState(fromId, "read_article_after_OIR_win")
			// srv.Db.EditStep(fromId, "+25.000₽ уходят в твой банк за правильный ответ!")
			// srv.SendMsgToServer(fromId, "bot", "+25.000₽ уходят в твой банк за правильный ответ!")

			// srv.SendAnimArticleHTMLV3("2.1", fromId, 2000)
			// srv.CopyMessage(fromId, -1001998413789, 25)
			// srv.SendAnimArticleHTMLV3("2.2", fromId, 2000)
			// srv.CopyMessage(fromId, -1001998413789, 27)

			// text := "тут должен быть вопрос"
			// reply_markup := `{"inline_keyboard" : [
			// 	[{ "text": "продолжить", "callback_data": "prodolzit_7_" }]
			// ]}`
			// srv.SendMessageWRM(fromId, text, reply_markup)

			
		} else {
			srv.SendMessage(fromId, "❌ Почта не найдена")
		}
	}

	if user.BotState == "read_article_after_KNB_win_2" {
		if !strings.HasPrefix(strings.ToLower(msgText), "хач") && !strings.HasPrefix(strings.ToLower(msgText), "хоч") {
			srv.SendMessageAndDb(fromId, "❌ Вы неверно ввели кодовое слово, сверьтесь с лонгридом и попробуйте еще раз")
			return nil
		}

		srv.SendAnimMessage("-1", fromId, animTimeout250)
		srv.SendBalance(fromId, "30.000", animTimeout250)
		srv.Db.EditStep(fromId, "9")
		srv.SendAnimMessageHTML("9", fromId, animTimeoutTest)


		// srv.ShowMilQ(fromId, 2)
		// srv.Db.EditStep(fromId, "7")
		srv.SendMsgToServer(fromId, "bot", msgText)
		return nil
	}

	if user.BotState == "read_article_after_OIR_win" {
		if !strings.HasPrefix(strings.ToLower(msgText), "рез") && !strings.HasPrefix(strings.ToLower(msgText), "риз") {
			srv.SendMessageAndDb(fromId, "❌ Вы неверно ввели кодовое слово, сверьтесь с лонгридом и попробуйте еще раз")
			return nil
		}

		// text := "Предлагаю тебе ответить на один вопрос 😏\nЗа него ты получишь +25.000₽ к банку💸"
		// replyMarkup := `{"inline_keyboard" : [
		// 	[{ "text": "Ествественно! Погнали!", "callback_data": "show_q_3_" }]
		// ]}`
		// srv.SendMessageWRM(fromId, text, replyMarkup)
		// srv.Db.EditStep(fromId, text)
		srv.Db.EditBotState(fromId, "")
		srv.SendAnimMessage("-1", fromId, animTimeout250)
		srv.SendBalance(fromId, "55.000", animTimeoutTest)
		srv.SendAnimMessageHTML("11", fromId, animTimeoutTest)
		srv.Db.EditStep(fromId, "11")
		srv.SendMsgToServer(fromId, "bot", "11 шаг")

		srv.ShowMilQ(fromId, 4)

		return nil
	}

	if user.BotState == "read_article_after_TrurOrFalse_win" {
		if !strings.HasPrefix(strings.ToLower(msgText), "син") {
			srv.SendMessageAndDb(fromId, "❌ Вы неверно ввели кодовое слово, сверьтесь с лонгридом и попробуйте еще раз")
			return nil
		}

		srv.Db.EditBotState(fromId, "")
		srv.SendBalance(fromId, "100.000", animTimeoutTest)
		srv.SendAnimMessageHTML("13", fromId, animTimeoutTest)
		srv.Db.EditStep(fromId, "13")
		time.Sleep(time.Second)

		text :=  "Если ты изучил всю информацию, то ты прямо сейчас можешь обменять свою награду 🏦 на способ заработка, который принесет тебе более 500.000₽ чистыми за раз 💸\n\nПлатить мне вперед не нужно, прибыль поделим пополам. Но поторопись, если хочешь вытащить прибыль несколько раз, ведь скоро способ перестанет работать. Жми кнопку ниже ⬇️"
		replyMarkup := `{"inline_keyboard" : [
			[{ "text": "Забрать схему", "url": "https://t.me/threeprocentsclub_bot" }]
		]}`
		srv.SendMessageWRM(fromId, text, replyMarkup)
		// srv.Db.EditStep(fromId, text)
		srv.SendMsgToServer(fromId, "bot", text)
		srv.Db.EditLatsActiontime(fromId, "")
		srv.Db.EditIsFinal(fromId, 1)
		
		return nil
	}

	return nil
}

func (srv *TgService) M_admin(m models.Update) error {
	fromId := m.Message.Chat.Id
	msgText := m.Message.Text
	fromUsername := m.Message.From.UserName
	srv.l.Info(fmt.Sprintf("M_start: fromId: %d, fromUsername: %s, msgText: %s", fromId, fromUsername, msgText))

	u, err := srv.Db.GetUserById(fromId)
	if err != nil {
		return fmt.Errorf("M_start GetUserById err: %v", err)
	}
	if u.Id == 0 {
		srv.SendMessage(fromId, "Нажмите сначала /start")
	}
	if u.IsAdmin != 1 {
		return fmt.Errorf("_______")
	}
	err = srv.ShowAdminPanel(fromId)

	return err
}
