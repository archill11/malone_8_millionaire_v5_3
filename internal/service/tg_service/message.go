package tg_service

import (
	"fmt"
	"myapp/internal/models"
	"myapp/pkg/files"
	my_regex "myapp/pkg/regex"
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
		if !strings.HasPrefix(strings.ToLower(msgText), "гоу") && !strings.HasPrefix(strings.ToLower(msgText), "го") && !strings.HasPrefix(strings.ToLower(msgText), "ко") && !strings.HasPrefix(strings.ToLower(msgText), "коу") && !strings.HasPrefix(strings.ToLower(msgText), "гэу") && !strings.HasPrefix(strings.ToLower(msgText), "go") {
			srv.SendMessageAndDb(fromId, "❌ Вы неверно ввели кодовое слово, сверьтесь с лонгридом и попробуйте еще раз")
			return nil
		}

		srv.Db.EditBotState(fromId, "")
		// srv.SendAnimMessage("-1", fromId, animTimeout250)
		// srv.SendBalance(fromId, "30.000", animTimeout250)
		// srv.Db.EditStep(fromId, "9")
		// srv.SendAnimMessageHTML("9", fromId, animTimeoutTest)

		text := "Ну что, поехали, ответь правильно на 3 вопроса и уже сегодня сможешь заработать 500.000₽ 😏"
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
