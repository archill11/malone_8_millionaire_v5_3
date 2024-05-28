package tg_service

import (
	"fmt"
	"time"
)

func (srv *TgService) ShowMilQ(chatId, qNum int) error {
	time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))

	textMap := map[int]string{
		1: "Первый вопрос 👆\n\nВыбери правильный ответ 👇",
		2: "Второй вопрос 👆\n\nВыбери правильный ответ 👇",
		3: "Третий вопрос 👆\n\nВыбери правильный ответ 👇",
		4: "Четвертый вопрос 👆\n\nВыбери правильный ответ 👇",
	}
	fileNameMap := map[int]string{
		1:  "./files/mil_q_1.jpg",
		2:  "./files/mil_q_2.jpg",
		3:  "./files/mil_q_3.jpg",
		4:  "./files/mil_q_4.jpg",
	}
	replyMarkupMap := map[int]string{
		1: `{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "_win_q_1_" }, { "text": "B", "callback_data": "_lose_q_1_" }, { "text": "C", "callback_data": "_lose_q_1_" }, { "text": "D", "callback_data": "_lose_q_1_" }]
		]}`,
		2: `{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "_lose_q_2_" }, { "text": "B", "callback_data": "_lose_q_2_" }, { "text": "C", "callback_data": "_win_q_2_" }, { "text": "D", "callback_data": "_lose_q_2_" }]
		]}`,
		3: `{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "_lose_q_3_" }, { "text": "B", "callback_data": "_lose_q_3_" }, { "text": "C", "callback_data": "_win_q_3_" }, { "text": "D", "callback_data": "_lose_q_3_" }]
		]}`,
		4: `{"inline_keyboard" : [
			[ { "text": "A", "callback_data": "_win_q_4_" }, { "text": "B", "callback_data": "_lose_q_4_" }, { "text": "C", "callback_data": "_lose_q_4_" }, { "text": "D", "callback_data": "_lose_q_4_" }]
		]}`,
	}

	text := textMap[qNum]
	replyMarkup := replyMarkupMap[qNum]
	fileNameInServer := fileNameMap[qNum]
	srv.SendPhotoWCaptionWRM(chatId, text, fileNameInServer, replyMarkup)

	srv.Db.EditStep(chatId, text)
	srv.SendMsgToServer(chatId, "bot", text)
	return nil
}

func (srv *TgService) Prodolzit(chatId int, prodolzit_id string) error {
	time.Sleep(time.Second * 2)

	if prodolzit_id == "1" {
		srv.SendAnimArticleHTMLV3("1.3", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 11)
		srv.SendAnimArticleHTMLV3("1.4", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 13)

		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_2_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		return nil
	}
	if prodolzit_id == "2" {
		srv.SendAnimArticleHTMLV3("1.5", chatId, 2000)
		
		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_3_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		return nil
	}
	if prodolzit_id == "3" {
		srv.CopyMessage(chatId, -1001998413789, 15)
		srv.SendAnimArticleHTMLV3("1.6", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 17)
		
		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_4_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		return nil
	}
	if prodolzit_id == "4" {
		srv.SendAnimArticleHTMLV3("1.7", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 19)
		
		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_5_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		return nil
	}
	if prodolzit_id == "5" {
		srv.SendAnimArticleHTMLV3("1.8", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 21)
		
		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_6_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		return nil
	}
	if prodolzit_id == "6" {
		srv.SendAnimArticleHTMLV3("1.9", chatId, 2000)
		return nil
	}
	if prodolzit_id == "7" {
		srv.SendAnimArticleHTMLV3("2.3", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 29)
		
		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_8_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		return nil
	}
	if prodolzit_id == "8" {
		srv.SendAnimArticleHTMLV3("2.4", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 31)
		
		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_9_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		return nil
	}
	if prodolzit_id == "9" {
		srv.SendAnimArticleHTMLV3("2.5", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 33)
		srv.SendAnimArticleHTMLV3("2.6", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 35)
		
		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_10_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		return nil
	}
	if prodolzit_id == "10" {
		srv.SendAnimArticleHTMLV3("2.7", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 37)
		srv.SendAnimArticleHTMLV3("2.8", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 39)
		
		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_11_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		return nil
	}
	if prodolzit_id == "11" {
		srv.SendAnimArticleHTMLV3("2.9", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 41)
		srv.SendAnimArticleHTMLV3("2.10", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 43)
		srv.SendAnimArticleHTMLV3("2.11", chatId, 2000)
		return nil
	}
	if prodolzit_id == "12" {
		srv.SendAnimArticleHTMLV3("3.3", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 51)
		
		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_13_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		return nil
	}
	if prodolzit_id == "13" {
		srv.SendAnimArticleHTMLV3("3.4", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 53)
		srv.SendAnimArticleHTMLV3("3.5", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 55)
		
		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_14_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		return nil
	}
	if prodolzit_id == "14" {
		srv.SendAnimArticleHTMLV3("3.6", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 57)
		srv.SendAnimArticleHTMLV3("3.7", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 59)
		srv.SendAnimArticleHTMLV3("3.8", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 65)
		srv.SendAnimArticleHTMLV3("3.9", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 66)
		
		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_15_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		
		return nil
	}
	if prodolzit_id == "15" {
		srv.SendAnimArticleHTMLV3("3.10", chatId, 2000)

		return nil
	}

	return nil
}

func (srv *TgService) ShowQWin(chatId int, q_num string) error {
	time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))
	time.Sleep(time.Second * 2)
	
	// textMap := map[string]string{
	// 	// "1":  fmt.Sprintf("+10.000₽ уходят в твой банк за правильный ответ!💸\n\nЧтобы разблокировать награду и забрать её - подпишись на мой канал 👇🏻\n\n%s\n\nИ жми кнопку ниже ⬇️", srv.Cfg.ChatLinkToCheck),
	// 	// "2":  "Снова в цель! ✅\nЕще один вопросик для победы 😏",
	// 	// "3":  "Прими мои поздравления, ты Снова дал верный ответ! 🎉🎉🎉",
	// 	// "4":  "А ты неплох 😏\nПравильный ответ ✅\n\nПереходим к следующему вопросу🔽",
	// 	"4":  fmt.Sprintf("+45.000₽ уходят в твой банк за правильный ответ!💸\n\nЧтобы разблокировать награду и забрать её - подпишись на мой канал 👇🏻\n\n%s\n\nИ жми кнопку ниже ⬇️", srv.Cfg.ChatLinkToCheck),
	// }

	if q_num == "1" {
		srv.Db.EditStep(chatId, "8")
		// srv.SendAnimMessageHTML("8", chatId, 2000)
		// text := "+19.000₽ уходят в твой банк за правильный ответ!💸\n\n🔐Чтобы разблокировать и забрать награду пришли мне кодовое слово из видео ☝🏻\n\n*Просмотр не займет много времени\nПосле пиши кодовое слово сюда.\nБуду ждать 👇🏻"
		// srv.SendVideoWCaption(chatId, text, "./files/VID_cod_1.mp4")
		// srv.CopyMessage(chatId, -1002074025173, 30) // https://t.me/c/2074025173/30

		srv.SendAnimMessageHTML("8", chatId, 2000)
		time.Sleep(time.Second * 2)
		text := fmt.Sprintf("Чтобы разблокировать награду и забрать её - подпишись на мой канал 👇🏻\n\n%s\n\nИ жми кнопку ниже ⬇️", srv.Cfg.ChatLinkToCheck)
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "Подписался☑️", "callback_data": "subscribe" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		srv.Db.EditStep(chatId, "Ссылка на канал")
		srv.SendMsgToServer(chatId, "bot", "Ссылка на канал")

		return nil
	}
	if q_num == "2" {
		srv.Db.EditStep(chatId, "10")
		// srv.SendAnimMessageHTML("10", chatId, 2000)
		// text := "+25.000₽ уходят в твой банк за правильный ответ!💸\n\n🔐Чтобы разблокировать и забрать награду пришли мне кодовое слово из видео ☝🏻\n\n*Просмотр не займет много времени\nПосле пиши кодовое слово сюда.\nБуду ждать 👇🏻"
		// srv.SendVideoWCaption(chatId, text, "./files/VID_cod_2.mp4")
		// srv.CopyMessage(chatId, -1002074025173, 31)
		srv.SendAnimMessageHTML("10", chatId, 2000)
		srv.Db.EditBotState(chatId, "read_article_after_KNB_win")
		// srv.Db.EditBotState(chatId, "read_article_after_OIR_win")
		srv.Db.EditStep(chatId, "+19.000₽ уходят в твой банк за правильный ответ!")
		srv.SendMsgToServer(chatId, "bot", "+19.000₽ уходят в твой банк за правильный ответ!")

		srv.SendAnimArticleHTMLV3("1.1", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 4) // https://t.me/c/1998413789/4
		srv.SendAnimArticleHTMLV3("1.2", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 9)

		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_1_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)
		
		// srv.SendAnimArticleHTMLV3("1.3", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 11)
		// srv.SendAnimArticleHTMLV3("1.4", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 13)
		// srv.SendAnimArticleHTMLV3("1.5", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 15)
		// srv.SendAnimArticleHTMLV3("1.6", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 17)
		// srv.SendAnimArticleHTMLV3("1.7", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 19)
		// srv.SendAnimArticleHTMLV3("1.8", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 21)
		// srv.SendAnimArticleHTMLV3("1.9", chatId, 2000)

		return nil
	}
	if q_num == "3" {
		srv.Db.EditStep(chatId, "12")
		// srv.SendAnimMessageHTML("12", chatId, 2000)
		// text := "+45.000₽ уходят в твой банк за правильный ответ!💸\n\n🔐Чтобы разблокировать и забрать награду пришли мне кодовое слово из видео ☝🏻\n\n*Просмотр не займет много времени\nПосле пиши кодовое слово сюда.\nБуду ждать 👇🏻"
		// srv.SendVideoWCaption(chatId, text, "./files/VID_cod_1.mp4")
		// srv.CopyMessage(chatId, -1002074025173, 32)
		srv.SendAnimMessageHTML("12", chatId, 2000)
		// srv.Db.EditBotState(chatId, "read_article_after_TrurOrFalse_win")
		srv.Db.EditBotState(chatId, "read_article_after_OIR_win")
		srv.Db.EditStep(chatId, "+25.000₽ уходят в твой банк за правильный ответ!")
		srv.SendMsgToServer(chatId, "bot", "+25.000₽ уходят в твой банк за правильный ответ!")

		srv.SendAnimArticleHTMLV3("2.1", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 25)
		srv.SendAnimArticleHTMLV3("2.2", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 27)

		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_7_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)

		// srv.SendAnimArticleHTMLV3("2.3", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 29)
		// srv.SendAnimArticleHTMLV3("2.4", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 31)
		// srv.SendAnimArticleHTMLV3("2.5", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 33)
		// srv.SendAnimArticleHTMLV3("2.6", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 35)
		// srv.SendAnimArticleHTMLV3("2.7", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 37)
		// srv.SendAnimArticleHTMLV3("2.8", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 39)
		// srv.SendAnimArticleHTMLV3("2.9", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 41)
		// srv.SendAnimArticleHTMLV3("2.10", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 43)
		// srv.SendAnimArticleHTMLV3("2.11", chatId, 2000)

		return nil
	}
	if q_num == "4" {
		srv.SendAnimMessageHTML("14", chatId, animTimeoutTest)
		srv.Db.EditStep(chatId, "14")
		srv.Db.EditBotState(chatId, "read_article_after_TrurOrFalse_win")

		srv.SendAnimArticleHTMLV3("3.1", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 46)
		srv.SendAnimArticleHTMLV3("3.2", chatId, 2000)
		srv.CopyMessage(chatId, -1001998413789, 49)

		text := "тут должен быть вопрос"
		reply_markup := `{"inline_keyboard" : [
			[{ "text": "продолжить", "callback_data": "prodolzit_12_" }]
		]}`
		srv.SendMessageWRM(chatId, text, reply_markup)


		// srv.SendAnimArticleHTMLV3("3.3", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 51)
		// srv.SendAnimArticleHTMLV3("3.4", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 53)
		// srv.SendAnimArticleHTMLV3("3.5", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 55)
		// srv.SendAnimArticleHTMLV3("3.6", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 57)
		// srv.SendAnimArticleHTMLV3("3.7", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 59)
		// srv.SendAnimArticleHTMLV3("3.8", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 65)
		// srv.SendAnimArticleHTMLV3("3.9", chatId, 2000)
		// srv.CopyMessage(chatId, -1001998413789, 66)
		// srv.SendAnimArticleHTMLV3("3.10", chatId, 2000)
		
		return nil
	}

	return nil
}

func (srv *TgService) ShowQLose(chatId int, q_num string) error {
	srv.ShowQLosePhoto(chatId, q_num)
	time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))

	// text := "Ответ неверный ❌\nК сожалению, ты ошибся, но шанс еще есть!\n\nЖми на кнопку 👇"
	// reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
	// 	[{ "text": "Попробовать еще раз", "callback_data": "show_q_%s_" }]
	// ]}`, q_num)
	// srv.SendMessageWRM(chatId, text, reply_markup)

	// srv.SendMsgToServer(chatId, "bot", text)
	return nil
}

func (srv *TgService) ShowQLosePhoto(chatId int, q_num string) error {
	time.Sleep(time.Millisecond * time.Duration(animTimeoutTest))

	user, err := srv.Db.GetUserById(chatId)
	if err != nil {
		return fmt.Errorf("ShowQLosePhoto GetUserById err: %v", err)
	}
	if user.Lives == 0 {
		return fmt.Errorf("0 жизней")
	}
	newLivesCnt := user.Lives - 1
	srv.Db.EditLives(chatId, newLivesCnt)

	pushTextMap := map[int]string{
		1: "❤️❤️🖤\nОтвет неверный ❌\n\nУ тебя сгорела одна жизнь 😔",
		2: "❤️🖤🖤\nОтвет неверный ❌\n\nУ тебя сгорела вторая жизнь, и это очень печально 😒",
		3: "🖤🖤🖤\nОтвет неверный ❌\n\nУ тебя сгорели все жизни 🥶\n\nНо у тебя еще есть шанс восстановить их.",
	}
	messIndex := 3 - newLivesCnt
	text := pushTextMap[messIndex]
	fileNameInServer := fmt.Sprintf("./files/push_%d.jpg", messIndex)
	if newLivesCnt != 0 {
		reply_markup := fmt.Sprintf(`{"inline_keyboard" : [
			[{ "text": "Попробовать еще раз", "callback_data": "show_q_%s_" }]
		]}`, q_num)
		_, err = srv.SendPhotoWCaptionWRM(chatId, text, fileNameInServer, reply_markup)
		if err != nil {
			return fmt.Errorf("ShowQLosePhoto SendPhotoWCaptionWRM err: %v", err)
		}
	} else {
		_, err = srv.SendPhotoWCaption(chatId, text, fileNameInServer)
		if err != nil {
			return fmt.Errorf("ShowQLosePhoto SendPhotoWCaptionWRM err: %v", err)
		}
	}
	srv.SendMsgToServer(chatId, "bot", text)

	if newLivesCnt == 0 && user.IsLastPush == 0 {
		huersStr, _ := srv.GetUserLeftTime(chatId)
		text = fmt.Sprintf("❗️У тебя есть %s на то, чтобы начать игру заново♻️\n\nЕсли ты не успеешь запустить игру за это время, то доступ к боту будет закрыт навсегда. Перезапуск бота не поможет, он просто перестанет работать для тебя ⛔️", huersStr)
		replyMarkup := `{"inline_keyboard" : [
			[{ "text": "ЗАБРАТЬ 100.000₽", "callback_data": "restart_game" }]
		]}`
		srv.SendMessageWRM(chatId, text, replyMarkup)
		srv.SendMsgToServer(chatId, "bot", text)
	}

	return nil
}
