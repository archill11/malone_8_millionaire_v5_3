package tg_service

import (
	"fmt"
	"myapp/internal/models"
	my_regex "myapp/pkg/regex"
	"strconv"
	"strings"
	"time"
)

func (srv *TgService) HandleCallbackQuery(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("HandleCallbackQuery: fromId: %d, fromUsername: %s, cq.Data: %s", fromId, fromUsername, cq.Data))

	srv.SendMsgToServer(fromId, "user", fmt.Sprintf("кнопка: %s", cq.Data))

	go func() {
		if cq.Data != "subscribe" {
			// time.Sleep(time.Second)
			srv.EditMessageReplyMarkup(fromId, cq.Message.MessageId)
			srv.Db.UpdateLatsActiontime(fromId)
		}
	}()

	// user, err := srv.Db.GetUserById(fromId)
	// if err != nil {
	// 	return fmt.Errorf("HandleCallbackQuery GetUserById err: %v", err)
	// }
	// if user.Id != 0 && user.Lives == 0 {
	// 	return nil
	// }

	if cq.Data == "delete_user_by_username_btn" {
		err := srv.CQ_delete_user_by_username_btn(m)
		if err != nil {
			srv.SendMessage(fromId, ERR_MSG)
			srv.SendMessage(fromId, err.Error())
		}
		return err
	}

	if cq.Data == "delete_user_by_id_btn" {
		err := srv.CQ_delete_user_by_id_btn(m)
		if err != nil {
			srv.SendMessage(fromId, ERR_MSG)
			srv.SendMessage(fromId, err.Error())
		}
		return err
	}

	if cq.Data == "start_game" {
		err := srv.CQ_start_game(m)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if cq.Data == "restart_game" {
		err := srv.CQ_restart_game(m)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if cq.Data == "subscribe" {
		err := srv.CQ_subscribe(m)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if strings.HasPrefix(cq.Data, "show_q_") { // показать mil вопрос
		if strings.Contains(strings.ToLower(cq.Message.Text), "ответ неверный") || (cq.Message.Caption != nil &&  strings.Contains(strings.ToLower(*cq.Message.Caption), "ответ неверный")) {
			time.Sleep(time.Second)
			srv.DeleteMessage(fromId, cq.Message.MessageId)
			srv.DeleteMessage(fromId, cq.Message.MessageId-1)
		}

		qId := my_regex.GetStringInBetween(cq.Data, "show_q_", "_")
		qIdInt, _ := strconv.Atoi(qId)
		err := srv.ShowMilQ(fromId, qIdInt)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if strings.HasPrefix(cq.Data, "_lose_q_") { // показать "Попробовать еще раз" на вопрос
		qId := my_regex.GetStringInBetween(cq.Data, "_lose_q_", "_")
		err := srv.ShowQLose(fromId, qId)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if strings.HasPrefix(cq.Data, "_win_q_") {
		qId := my_regex.GetStringInBetween(cq.Data, "_win_q_", "_")
		err := srv.ShowQWin(fromId, qId)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if strings.HasPrefix(cq.Data, "prodolzit_") { // prodolzit_14_
		prodolzit_id := my_regex.GetStringInBetween(cq.Data, "prodolzit_", "_")
		err := srv.Prodolzit(fromId, prodolzit_id)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	if cq.Data == "mailing_copy_btn" {
		err := srv.CQ_mailing_copy_btn(m)
		if err != nil {
			srv.SendMessageAndDb(fromId, ERR_MSG)
			srv.SendMessageAndDb(fromId, err.Error())
		}
		srv.Db.UpdateLatsActiontime(fromId)
		return err
	}

	srv.Db.UpdateLatsActiontime(fromId)
	return nil
}

func (srv *TgService) CQ_start_game(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("CQ_start_game: fromId: %d, fromUsername: %s", fromId, fromUsername))

	srv.SendAnimMessage("-1", fromId, animTimeout250)
	srv.SendBalance(fromId, "1000", animTimeout250)
	srv.SendAnimMessageHTML("2", fromId, animTimeoutTest)
	srv.SendAnimMessage("4", fromId, animTimeoutTest)
	srv.Db.EditStep(fromId, "5")
	srv.SendAnimMessage("5", fromId, animTimeoutTest)

	err := srv.ShowMilQ(fromId, 1)
	if err != nil {
		return fmt.Errorf("CQ_start_game ShowMilQ1 err: %v", err)
	}

	return nil
}

func (srv *TgService) CQ_mailing_copy_btn(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("CQ_start_game: fromId: %d, fromUsername: %s", fromId, fromUsername))

	srv.SendForceReply(fromId, MAILING_COPY_STEP)

	return nil
}

func (srv *TgService) CQ_restart_game(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	fromFirstName := cq.From.FirstName
	srv.l.Info(fmt.Sprintf("CQ_restart_game: fromId: %d, fromUsername: %s", fromId, fromUsername))

	user, err := srv.Db.GetUserById(fromId)
	if err != nil {
		return fmt.Errorf("CQ_restart_game GetUserById err: %v", err)
	}
	if user.CreatedAt != "" && srv.IsIgnoreUser(fromId) {
		return nil
	}

	err = srv.Db.AddNewUser(fromId, fromUsername, fromFirstName)
	if err != nil {
		return fmt.Errorf("CQ_restart_game AddNewUser err: %v", err)
	}
	srv.Db.EditBotState(fromId, "")
	srv.Db.EditLives(fromId, 3)
	srv.SendMessageAndDb(fromId, fmt.Sprintf("Привет, %s 👋", fromFirstName))

	srv.Db.EditStep(fromId, "1")
	srv.SendAnimMessageHTML("1", fromId, animTimeout3000)

	time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))
	
	text := "Прямо сейчас начинай игру и забирай бонус 1000₽ за уверенный старт! 🚀"
	replyMarkup := `{"inline_keyboard" : [
		[{ "text": "Начать игру", "callback_data": "start_game" }]
	]}`
	srv.SendMessageWRM(fromId, text, replyMarkup)
	
	srv.SendMsgToServer(fromId, "bot", text)
	srv.Db.UpdateLatsActiontime(fromId)

	return nil
}

func (srv *TgService) CQ_subscribe(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("CQ_subscribe: fromId: %d, fromUsername: %s", fromId, fromUsername))

	GetChatMemberResp, err := srv.GetChatMember(fromId, srv.Cfg.ChatToCheck)
	if err != nil {
		return fmt.Errorf("CQ_subscribe GetChatMember fromId: %d, ChatToCheck: %d, err: %v", fromId, srv.Cfg.ChatToCheck, err)
	}
	if GetChatMemberResp.Result.Status != "member" && GetChatMemberResp.Result.Status != "creator" {
		logMess := fmt.Sprintf("CQ_subscribe GetChatMember bad resp: %+v", GetChatMemberResp)
		srv.l.Error(logMess)
		mess := "❌ вы не подписаны на канал!"
		srv.SendMessageAndDb(fromId, mess)
		srv.Db.EditStep(fromId, mess)
		return nil
	}

	go func() {
		time.Sleep(time.Second)
		srv.EditMessageReplyMarkup(fromId, cq.Message.MessageId)
	}()

	// srv.Db.EditBotState(fromId, "")
	srv.SendAnimMessage("-1", fromId, animTimeout250)
	// srv.SendBalance(fromId, "100.000", animTimeoutTest)
	// srv.SendAnimMessageHTML("13", fromId, animTimeoutTest)
	// srv.Db.EditStep(fromId, "13")
	srv.SendAnimMessage("-1", fromId, animTimeout250)
	srv.SendBalance(fromId, "11.000", animTimeout250)
	srv.Db.EditStep(fromId, "7")
	srv.SendAnimMessageHTML("7", fromId, animTimeoutTest)
	// srv.CopyMessage(fromId, -1002074025173, 22)
	time.Sleep(time.Second)

	text := "Предлагаю тебе ответить на один вопрос 😏\nЗа него ты получишь +19.000₽ к банку💸"
	replyMarkup :=`{"inline_keyboard" : [
		[ { "text": "Давай попробуем", "callback_data": "show_q_2_" } ]
	]}`
	srv.SendMessageWRM(fromId, text, replyMarkup)
	// srv.ShowMilQ(fromId, 2)
	srv.Db.EditStep(fromId, "7")
	srv.SendMsgToServer(fromId, "bot", "7 шаг")

	// text :=  "Если ты изучил всю информацию, то ты прямо сейчас можешь обменять свою награду 🏦 на способ заработка, который принесет тебе более 500.000₽ чистыми за раз 💸\n\nПлатить мне вперед не нужно, прибыль поделим пополам. Но поторопись, если хочешь вытащить прибыль несколько раз, ведь скоро способ перестанет работать. Жми кнопку ниже ⬇️"
	// replyMarkup := `{"inline_keyboard" : [
	// 	[{ "text": "Забрать схему", "url": "https://t.me/threeprocentsclub_bot" }]
	// ]}`
	// srv.SendMessageWRM(fromId, text, replyMarkup)
	// // srv.Db.EditStep(fromId, text)
	// srv.SendMsgToServer(fromId, "bot", text)
	// srv.Db.EditLatsActiontime(fromId, "")
	// srv.Db.EditIsFinal(fromId, 1)

	return nil
}

func (srv *TgService) CQ_delete_user_by_username_btn(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("CQ_delete_user_by_username_btn: fromId: %d, fromUsername: %s", fromId, fromUsername))

	srv.SendForceReply(fromId, DEL_USER_MSG)
	return nil
}

func (srv *TgService) CQ_delete_user_by_id_btn(m models.Update) error {
	cq := m.CallbackQuery
	fromId := cq.From.Id
	fromUsername := cq.From.UserName
	srv.l.Info(fmt.Sprintf("CQ_delete_user_by_id_btn: fromId: %d, fromUsername: %s", fromId, fromUsername))

	srv.SendForceReply(fromId, DEL_USER_ID_MSG)
	return nil
}
