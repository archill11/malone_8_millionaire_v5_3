package tg_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"myapp/internal/models"
	"myapp/internal/repository/pg"
	"myapp/pkg/logger"
	"myapp/pkg/my_time_parser"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)


var (
	mskLoc, _ = time.LoadLocation("Europe/Moscow")
)


var (
	articlesMap = map[string][]string{
		"1.1": {
			"Если ты здесь, то ты уже как минимум на 1 шаг дальше 70% людей запустивших игру \"Кто хочет стать миллионером\"• и только 30% из запустивших, остались и читают этот текст.🤷‍♂️• \n Да, это такая статистика• и буду честен, этот бот создан специально для отсеивания дурачков•(NPC, ботов, в общем называйте как хотите🤣),• которые даже на элементарные вопросы не могут ответить.• \nНичего личного, буду с вами откровенен, но такова правда жизни.• <u>Дураки сами обычно не доходят до больших денег</u> 🤡,• хотя возможности у них есть,• стоило просто попытаться ещё.• \nПример статистики игры:• \nК слову о том, почему каждый в мире не сможет стать богатым.• На скрине выше вы видите пример статистики доходимости аудитории в боте.• Из 1362 людей запустивших игру, смогли ответить на 6 вопросов• и дойти до этого текста всего 818 человек✅.• И чем дальше будет продолжаться игра,• тем больше людей будет тупо покидать её. \nЛюди сами отсеиваются,• никто не заставляет их прозябать в бедности.• Ну да ладно,• там дурачкам и место.•",
		},
		"1.2": {
			"Под дурачками я подразумеваю именно гиперленивых людей🦥,• а не интеллектуальные способности или образование.•\n\nНапример, я сам закончил только 9 классов,• но дефицит пятёрок в школе не помешал их профициту в моём кармане сейчас😏• Я просто не опустил руки в нужный момент и не сдался❕•\n\nДа, говорят, что лень - <i>двигатель прогресса.</i>• От части это так, но работает это на ряду с положительными качествами человека,• такими как, например, <u>целеустремлённость.</u>• Если кроме лени у тебя ничего нет,• то не льсти себе, ты всего лишь обычный диванный критик.•\n\n\nИтак!• Как же я буду отдавать тебе обещанные деньги, если ты пройдёшь игру до конца?• А всё просто!•😉•\n\nПрямо сейчас в мои руки попал инструмент,• благодаря которому можно <u>откусить значительную часть пирога🍰(очень значительную) от новой ниши,• пока сюда не зашла гора конкурентов!</u>•\n\nА это значит,• что одного моего рта на этот пирог недостаточно.•\n\nМне понадобится много времени,• которого у меня нет,• ведь конкуренты не спят• и каждый день сюда приходят голодные ребята🦁,• которые так же хотят больше,• чем могут проглотить.❕•",
		},
		"1.3": {
			"Изначально, я хотел делиться этим инструментом со всеми платно 💵,• чтобы <u>хоть что-то получить от дурачков,• которые просто так тратят моё время</u>,• но моя ненависть к инфоцыганскому поприщу всё-таки не позволила мне сделать это• и тогда я придумал эту игру🔥 .•\n\n\n И всё, что тебе нужно,• чтобы без каких-либо потерь набить свой карман деньгами уже сегодня -• это пройти эту игру до конца.• <b>Эта игра -• действительно панацея от бедности</b>💊💰.•\n\nА перед тем как я расскажу тебе о том,• что сделает тебя миллионером уже через пару-тройку недель,• давай немного познакомимся и я представлюсь:•\n\nВ общем, меня зовут Марк Одинцов,• мне 29 лет🖐.• Родом я с самой обычной глуши в России,• обычный региональный парниша.• Жили мы мягко говоря бедно,• но родители всегда старались одаривать меня любовью и заботой,• поэтому не смотря на не лучшие условия проживания,• воспоминания о детстве у меня очень тёплые❤️.•",
		},
		"1.4": {
			"<u>Дальше я завалил ЕГЭ</u> 🤯• (не знаю как этот экзамен сейчас называется)• и пошёл учиться в ПТУ.• Параллельно я начал работать официантом в местной забегаловке• и вот тогда то я понял,• что вкалывать за копейки не моё🙅‍♂️ . •\n\nВ таком общежитии я жил во время обучения в ПТУ,• я наконец начал понимать родителей,• которые постоянно находились на работе,• приходили домой уставшие и вымотанные• и еще находили силы заняться моим воспитанием и поиграть со мной. 👨‍👩‍👦❤️•\n\nЯ твёрдо решил• что не буду работать в найме• и буду заниматься своим делом.• \nМною дополнительно двигало чувство благодарности родителям.• Они всё делали для того,• чтобы я мог вкусно покушать,• ну а теперь и я должен им котлету)💰•\n\nНе буду растягивать рассказ о себе,• т.к. он в принципе похож на большинство тех, что вы наверняка уже слышали 🫣•\nПочему то, так действительно получается,• что 80% историй успеха начинаются с бедности,• поэтому воспользуйтесь этим, как преимуществом 💸•\n\nРедко выросший в тепличных условиях, может хотя бы сохранить успех родителей,• не говоря уже о своём собственном.😔•",
		},
		"1.5": {
			"Если вкратце,• с момента как я принял решение работать на себя,• прошло не мало времени, когда я ещё работал в найме. 🥵•\n\nПризнавайтесь!• Узнали себя в этом мемчике?😉•\nУ меня ничего не получалось 7 лет,• т.к. тогда не было таких обильных обучений и возможностей как сейчас.• \n\n<u>Я протаптывал всю дорогу себе сам,• без партнеров и наставников и вот сейчас я здесь</u>.• Надеюсь у вас не осталось сомнений,• что за эти годы я стал действительно компетентен в том, чем занимаюсь ❗️•\n\nСейчас я живу прекрасно,• у меня огромные аппартаменты в солнечном Дубае🏛,• личные ассистенты и активно набирающаяся команда из энтузиастов• и возможно именно ты пополнишь ряды моей команды,• если сможешь пройти эту игру😉•\nСейчас мои аппарты выглядят примерно так.• Это конечно не общажная движуха,• но тоже сойдет)😄•\n\nА теперь внимание ❗️•\nМне нужны люди,• которым я предоставлю <u>пошаговый алгоритм действий</u> • и после того как вы получите с этого деньги -• вернете мне 50%🤝.• \n\nВход в мою команду стоит 100.000₽• НО!• Пройдя игру до конца ты выиграешь как раз эту сумму• и сможешь потратить их на вход.•",
		},
		"1.6": {
			"<b>То есть, фактически, вход для победителей игры БЕСПЛАТНЫЙ</b> 🆓 • \n\nЯ уже не знаю, что можно ещё лучше предложить людям,• чтобы они начали наконец действовать.• Так выглядит 97%• населения земли🤣• (смотреть со звуком)•\nВ общем, по-моему справедливая сделка win-win🏆,• в которой <u>выигрывают обе стороны.</u>• \n\nВ этом и заключается моя выгода• и то для чего я сделал всё это• и теперь ты читаешь этот текст находясь в шаге от <b>финансовой свободы</b>💸.• \n\nВсё в белую,• легально и абсолютно законно,• никакой чернухи!• Обычная лазейка,• которую можно выжать, если не щелкать клювом🤫.• А всё что тебе нужно -• это потратить 15-20 минут и пройти игру \"Кто хочет стать миллионером\" до конца!•\n<u>Чтобы продолжить игру вводи в бота кодовое слово \"хочу\" </u>•\n\n\n<b>Что тебя ждёт дальше:</b>•\n🎁 +19.000₽ на баланс уже сейчас.•\n\n🎁 Мини-игра которая даст ещё +25.000₽ в случае победы. •\n\n🎁 Объяснения по направлению, которым мы будем заниматься и сколько можно с этого  заработать.•",
		},

		"2.1": {
			"|<Результаты и доказательства.>| 🔥•\n\nОбычно на этом этапе отклеивается еще 15% из тех,• кто запустил игру и остаётся так же 15%✅. •\n\n_<Людям просто лень потратить 5 минут на прочтение>_😁• Тем и лучше!• Это работает для меня как элементарный фильтр,• ведь расп#здяи мне не нужны,• пускай лучше бабки получат те,• кому это действительно необходимо,• кто этого заслуживает и кто готов е*ашить! •\n\nИ если ты читаешь этот текст,• то |<ты входишь в 15% оставшихся людей,>|• а это уже не мало• и поэтому я зову тебя с собой в увлекательное путешествие.• \nСнимем сливки и закрепимся навечно в этой нише!• Я это гарантирую🤝•\n\nНачну с того, что каждый год зарождается минимум 2 новых направления♻️,• которые дают |<бешенные денежные результаты>|• за счёт того, что \"пузырь раздувается\"• и оказавшись у истоков зарождения этой темы• можно вылететь из нищеты• как пробка из под шампанского🍾. •\n\n_<Первое время мы будем делать с вами до сотни иксов>_.• Да, это не опечатка и не автозамена,• мы будем приумножать капитал х80-100 ❗️. •\n\nСпустя какое-то время, когда это выжмется,• конверсия |<упадёт и будет давать уже десятки иксов(х20-х50),>|• на этом этапе мы будем продолжать выжимать мою технологию,• ведь всё равно это достаточно много💲.•",
		},
		"2.2": {
			"И только потом• |<рынок заполнится кучей зевак,>|• пузырь начнёт сдуваться активно• и давать уже всего несколько иксов 📉•\n\nНа этом этапе мы уже заработаем достаточно 💸• и я начну поиск другого пузыря. •\nСколько бы не занял поиск,•нам _<хватит заработанных денег• минимум на 10 безбедных лет вперёд>_🤑,• но мне обычно достаточно не более 1 года,• чтобы найти что-то новое. •\n\nВсегда появляется новый пузырь,• это лишь вопрос времени• и 1 год - это я беру с запасом,• чтобы вы не строили |<неоправданных надежд• и не разочаровывались.>| 😎•\n\nВы думаете откуда берутся все эти молодые долларовые миллионеры❓•\nЧто-то вкладывали❓•Пффф.• Упорный труд ❓•Я вас умоляю, •в их то годы... •|<Все они просто когда-то оказались в нужное время в нужном месте>|•😉.•\n\nНо сейчас давайте не будем загадывать наперёд,• а сосредоточимся на том, что имеем. •\nВы наверное спросите• \"Ну не может же быть всё так просто. В чём подвох?\"•🤔•",
		},
		"2.3": {
			"Я отвечу так: •\n\n1️⃣|<Во-первых:>|• приумножить можно не любой капитал💰•, а лишь небольшую часть с каждого человека. •_<Т.е., условно, с 1.000.000₽ вы не сможете сделать 100.000.000₽.>_ •Тогда бы мне не нужны были вы. •Приумножаются маленькие суммы• и именно поэтому мне нужно больше рук. •\n\n2️⃣|<Во-вторых:>| • один человек сможет работать по инструкции не чаще 1 раза в неделю, •но и это решаемо.• Я объясню вам как и что делать,• а дальше уже |<зависит от вашего энтузиазма>|.• Захотите - сделаете больше,• по сути ничего сложного, нужно только желание✊.•\n\n3️⃣|<В-третьих:>|• это для вас просто)• Для меня же это месяцы поисков и тестов,• чтобы получить готовый и рабочий инструмент,• который я предоставляю вам🔥.• И вот поэтому я хочу, чтобы вы откинули некоторые заблуждения:•\n\nЗаблуждение #1.• 《|<Сделаю сам, мне кажется я смогу заработать больше, если сделаю по своему>|》•\n\nПолучив от меня конкретную инструкцию,• нужно следовать строго по ней• и не применять никакой отсебятины❗️• \n\nВ какой-то момент вам может показаться,• что вы лучше меня знаете как действовать• и сможете больше заработать в моменте, если сделаете как-то по своему. •\n\nНи в коем случае не превращайтесь в |<азартных самонадеянных идиотов>|❗️•",
		},
		"2.4": {
			"Нам нужна стабильность• и следуя моей инструкции• вы в короткие сроки сможете |<заработать миллионы>|💵. •\nНе жадничайте• и только тогда будете щедро вознаграждены,• действуйте с холодной головой🤯•\n\nНадеюсь, все объяснил предельно понято. •\n\nРезюмируя - делай все по инструкции •и в этом случае ты с вероятностью 146% заработаешь на моей теме.• Согласен❓ •\n\nДавай перейдем к следующему заблуждению ⤵️•\n\nЗаблуждение #2• «|<Сейчас не готов. Отложу на потом, когда будет более подходящий момент>|»•\n\n_<Откладывай безделье, да не откладывай дела.>_•\nЯ всегда придерживаюсь одного принципа:• не важно сколько я зарабатываю:• 50 тысяч или 50 миллионов в месяц,• я всегда работаю одинаковое количество времени ⌛️•\n\nНасколько плоха или хороша не была бы моя жизнь• - всегда можно сделать лучше, •если ты просто не откладываешь,• а делаешь здесь и сейчас⏳. •\n\n_<Лично для меня нет более подходящего времени чем сейчас,>_• но давай отбросим всю эту лирику про пахоту 24/7• и достижение успеха потом и кровью 😅🩸•",
		},
		"2.5": {
			"Я тебе назову одну весомую причину,• по которой ты должен начать именно сейчас:• актуальность темы. 💸•\n\nДа, как я и говорил ранее,• в год появляется около 2 тем,• которые |<выстреливают и дают сумасшедшие деньги>|🤯• но так будет не долго,• нужно оказаться на зарождении этой ниши,• встать в ряды раньше всех❕•\n\nТолько в этом случае• у тебя получится сделать на этом весомые результаты ✅•\n_<Зайдя позже, вряд-ли получится не то что заработать, •но и еще очень вероятно, что потеряеш>_😔•\nДа и в принципе• откладывая что-то в долгий ящик,• оно остается там навсегда с вероятностью 99%,• это доказанный факт. 📦•\n\nА даже если ты решишь попробовать,• то может быть уже поздно;• мы то с командой уже подняли бабла• и |<чиллим где-нибудь в Амстердаме• или на Мальдивах на заработанные бабки>|😎. •\n\nЛегко может произойти ситуация, что тема выжата• и работать больше не будет, •а когда появится следующая подобная тема никто не знает 🤷‍♂️•",
		},
		"2.6": {
			"|<Заблуждение #3>| •\nНе бывает так, чтобы каждый мог зарабатывать много ⛔️•\n\nСамо выражение сложно назвать заблуждением.• Действительно все в мире не могут быть богатыми,• но ведь нам все и не нужны) 😎•\n\nНа предыдущей странице я рассказывал• о том для чего придумал бота \"|<Кто хочет стать миллионером>|\". •\nДля меня это отсев лишней аудитории👥•\nЯ наблюдал поведение людей запускающих бота:• до куда они доходили, дочитывали и т.д. •\n\n_<Вот вам наглядный пример, как выглядит статистика,• начиная с запуска бота и до его прохождения и входа в команду >_👇•\n\nЭто действительно такая статистика.• Из 1000 людей запустивших бота, •результат в итоге получат только 30✅.• |<Это 3%, что вполне логично,• ведь именно такой процент населения планеты имеет высокий доход>|. 💸💸•\nЯ не из головы это взял,• погуглите про это, если хотите🧠•\n\nТ.е. люди сами по себе отсеиваются на ровном месте, •поэтому не смотрите на количество подписчиков и охватов постов. •\n|<Обычно 97% из всего этого числа - зеваки>| 🥱• предпочитающие просто наблюдать, а потом ворчать на чужие успехи,• оправдывая своё бездействие.•",
		},
		"2.7": {
			"Я прогорал тысячу раз•, оказавшись либо на закате какой-то темы📈,• либо доверившись не тем людям😔•\n\nИ на данный момент я могу уверенно сказать, что \"съел собаку\"• на поиске денежного направления• и вижу какие-то |<сочные ниши еще на этапе их зарождения>|💰. •\n\nИ как раз сейчас я заметил такую!• Я протестировал её и могу сказать однозначно:• _<если следовать чётко по моей инструкции, то профит неизбежен>_,• никаких рисков. Быстрые белые деньги❗️•\n\nЭто как нефтяная скважина• из которой вот-вот фонтаном начнёт бить нефть• и нам нужно просто успеть подставить как можно больше ведер и быть заранее готовыми,• чтобы унести с рынка значительную часть🤑.•\nНу и на этой части текста осталось всего 10% •дочитавших из всех кто начал читать. •\n\nА значит я готов приоткрыть вам завесу самого сладкого🍯:• |<сколько денег на этом зарабатывают сейчас люди>|, которых я подтянул из своего окружения •и сколько на этом сможешь заработать ты уже сегодня❗️•\n\nСейчас скажу только одно: •эта _<технология была протестирована лично мной и моей командой на нескольких проектах>_✅•",
		},
		"2.8": {
			"|<И везде она изумительно себя показала>|.• Ни одного промаха. •Вот что значит сила зарождающегося пузыря!✊•\n\n|<Начну со своих результатов>|. •\nНа скринах ниже вы видите мой баланс до того, как я вошёл в эту тему• и спустя неделю👇•\n\nА вот такие чеки можешь получать ты 1-2 раза в неделю,• если сможешь пройти игру до конца 👇•\n\n_<Чтобы продолжить игру вводи в бота кодовое слово \"результат\">_. •\n\nНо учти... То, о чём я поведаю тебе дальше — сильно противоречит информации,• которую ты слышал от многочисленных “гуру”.•\n\n❗️ Это действительно |<малоизвестная информация>|. •И ты в одном шаге от ознакомления с ней.•\n\nКак это может поменять жизнь любого человека❓•",
		},

		"3.1": {
			"_<Вот мы и подобрались к заключающей части игры.>_• На этом этапе остаётся лишь 6% людей✅,• начавших игру, а значит именно вы получите подарок стоимостью 100.000₽🎁,• только дочитайте внимательно до конца.•\n\nНа данный момент я |<активно набираю людей команду,>|• в которой мы будем нагло выкачивать бабки из надувающегося пузыря• и как я писал ранее -• снимем самые жирные сливки💸.•",
		},
		"3.2": {
			"Когда я понял,• что без команды заберу с рынка лишь маленькую крупицу📉,• я начал искать людей и как поступил бы, думаю, любой из вас• - я начал предлагать своим родным и близким людям |<стать частью команды>|😬. •\n\nКак бы прискорбно не было осознавать это,• но даже если среди твоих родных и близких есть |<слабовольные и тупые люди,>|• ты при всём желании не сможешь помочь им,• даже если всё будешь делать за них😞.•\n\n_<Чисто физически тебе не хватит рук пахать за десятерых.>_• Да и зачем помогать даже близкому человеку,• которому самому по барабану на своё финансовое состояние🤷‍♂️.•  \n\nТогда то я и начал думать, что мне нужен какой-то инструмент,• который сам будет отсеивать• способных людей от неспособных. 📝•\n\nПосле долгих разработок и тестов появилась игра \"|<Кто хочет стать миллионером>|\",• благодаря которой ты находишься здесь• буквально в сантиметре от новой жизни• без долгов, кредитов,• ипотек и бытовухи🤩💸•",
		},
		"3.3": {
			"Сейчас я хочу вас познакомить с первым человеком,• который попал в команду пройдя игру.• |<Знакомьтесь - Алина>|😄•\n\nКогда Алина прошла игру и написала мне, чтобы вступить в команду,• я уже начал думать, что моя игра всё таки не работает •и не отсеивает ненужных мне людей.• Ведь когда я начал узнавать как она живёт,• мне показалось \"|<Ну не, она точно не потянет>|\"🤦‍♂️•\n\nЧтобы не растягивать статью -• старался уместить переписку на один скрин,• поэтому приближай и читай)💬•\n\n_<Не смотря на свой скептицизм,• я всё же решил скинуть ей инструкцию>_.• Сказать честно -• я думал она пропадёт,• перестанет писать мне, отвечать,• ну или максимум придумает какую-нибудь отговорку,• лишь бы ничего не делать.😬•\n\n|<Но каково было моё удивление,• когда спустя пару часов она отписала мне с результатом>|.•\nТогда то я и понял, что никогда ещё так не ошибался в людях.•\n\nВот, Алина пишет мне в шоке с того, что всё получилось😉.• Честно говоря, я был удивлён не меньше неё, •но старался не подавать виду.• |<Увеличивай скрины и читай>|•\n\nПотом мой неугомонный здравый смысл• начал искать какие-либо причины снова сомневаться в Алине.😦•",
		},
		"3.4": {
			"Я думал:• \"Скорее всего она просто сделала как надо один раз, а дальше будет косячить и пропадать\".• |<Но и тут я очень сильно ошибся>|.😨•\n\nКак видите на скринах ниже, спустя более чем 2 месяца,• мы до сих пор сотрудничаем• и за это время Алина ни разу не подвела меня✊. •\n\nЗато за это время она:• _<погасила ипотеку, закрыла кредиты, заработала на отдых и уже копит на новую квартиру>_ • и уверен копить она будет совсем недолго😉.• \n\n|<Увеличивайте и читайте как изменилась её жизнь• всего за несколько месяцев>|💸•\n\nКакой вывод я могу сделать из ситуации с Алиной:• |<не важно в какой жизненной ситуации вы находитесь сейчас• и в каких условиях живёте>| 😔. •\nЕсли вы обладаете необходимыми качествами -• |<вы обречены на успех>|❗️•\n\n _<Просто у вас нет стартового капитала, либо нужного окружения единомышленников>_,• либо необходимых связей. •\nИли всего этого вместе, не важно, •но у вас точно есть определенный склад ума, который мне нужен,• |<чтобы мы могли вместе свернуть горы>|🔥.•\n\nДоверьтесь мне• и мы поможем друг другу сколотить денежное состояние,• обещаю!🤝💰•",
		},
		"3.5": {
			"|<Если человек способен изменить свою жизнь к лучшему,• ему нужно только дать возможность и он её не упустит.>| •\n\nВцепится зубами в неё• и не в коем случае _<не разожмёт челюсть , пока не получит результат.>_ •\n\nИ сейчас я готов дать такую |<возможность тебе>|❗️•\nДавайте представим, что я Морфеус из фильма \"Матрица\"• и предлагаю вам 2 капсулы на выбор💊. •\n\n|<Первая красная>|💔,• которая продолжит убивать вас скучным бытом и серой невзрачной жизнью,• в которой вы не можете ничего позволить ни себе, ни своим родным😔:•\n\nТвоя жизнь после выбора красной капсулы 👆•\n\n|<А вторая синяя>|💙,• которая пробудит вас и перевернёт вашу жизнь с ног на голову• в самом прекрасном смысле 😄: •\n\nТвоя жизнь после выбора синей капсулы 👆•\n\n\nКазалось бы, выбор прост,• но 97% людей выбирают красную капсулу💔. •\n_<Какую выберешь ты, решать только тебе>_:•\n\nСделай правильный выбор!•\nЕсли ты хочешь пробудиться и выбираешь синюю капсулу, чтобы•",
		},
		"3.6": {
			"Кардинально поменять свою\nЖизнь в лучшую сторону ,• то вводи в бота кодовое слово «СИНЯЯ» •\n\n|<Что тебя ждёт дальше>|❓•\n\nТы наконец-то пройдешь игру и получишь мой контакт• по которому сможешь написать мне в ЛС 🔥. •\n\nМы с тобой познакомимся,• я скину тебе инструкцию• и уже _<сегодня ты заработаешь первые деньги>_• и создашь себе стабильный источник дохода💸. •\n\nДерзай! •Игра можно сказать пройдена, жду тебя в ЛС 💬•",
		},
	}

	stepsMap = map[string][]string{
		"-1": {"🏦⁣  💰💰💰👈", "🏦⁣ 💰💰💰👈", "🏦⁣💰💰💰👈", "🏦⁣💰💰👈", "🏦⁣💰👈", "🏦⁣👈", "delete"},

		"1": {
			"Я предлагаю тебе сыграть со мной в игру,",
			"Я предлагаю тебе сыграть со мной в игру, где ты будешь получать деньги за каждое правильно выполненное задание 💸",
			"Я предлагаю тебе сыграть со мной в игру, где ты будешь получать деньги за каждое правильно выполненное задание 💸\n\nЕсли пройдешь игру до конца, <b>то сможешь выиграть 100.000 рублей 😳</b>",
			"Я предлагаю тебе сыграть со мной в игру, где ты будешь получать деньги за каждое правильно выполненное задание 💸\n\nЕсли пройдешь игру до конца, <b>то сможешь выиграть 100.000 рублей 😳</b>\nВся игра не займёт более 15 минут ⌛️",
		},
		"2": {
			"<b>Кстати, у тебя есть 3 жизни ❤️</b>",
			"<b>Кстати, у тебя есть 3 жизни ❤️</b>\n\n🥵Если ты долго бездействуешь, то сгорает одна жизнь.",
			"<b>Кстати, у тебя есть 3 жизни ❤️</b>\n\n🥵Если ты долго бездействуешь, то сгорает одна жизнь.\n😳 Дашь неверный ответ на вопрос - сгорает одна жизнь",
			"<b>Кстати, у тебя есть 3 жизни ❤️</b>\n\n🥵Если ты долго бездействуешь, то сгорает одна жизнь.\n😳 Дашь неверный ответ на вопрос - сгорает одна жизнь\n😔Когда сгорают все 3 жизни, твой баланс обнуляется и ты проигрываешь эту игру",
			"<b>Кстати, у тебя есть 3 жизни ❤️</b>\n\n🥵Если ты долго бездействуешь, то сгорает одна жизнь.\n😳 Дашь неверный ответ на вопрос - сгорает одна жизнь\n😔Когда сгорают все 3 жизни, твой баланс обнуляется и ты проигрываешь эту игру\n\nУ тебя будет немного времени, чтобы начать всё заново,",
			"<b>Кстати, у тебя есть 3 жизни ❤️</b>\n\n🥵Если ты долго бездействуешь, то сгорает одна жизнь.\n😳 Дашь неверный ответ на вопрос - сгорает одна жизнь\n😔Когда сгорают все 3 жизни, твой баланс обнуляется и ты проигрываешь эту игру\n\nУ тебя будет немного времени, чтобы начать всё заново, но лучше пройди всё с первой попытки🏆",
			"<b>Кстати, у тебя есть 3 жизни ❤️</b>\n\n🥵Если ты долго бездействуешь, то сгорает одна жизнь.\n😳 Дашь неверный ответ на вопрос - сгорает одна жизнь\n😔Когда сгорают все 3 жизни, твой баланс обнуляется и ты проигрываешь эту игру\n\nУ тебя будет немного времени, чтобы начать всё заново, но лучше пройди всё с первой попытки🏆\n\nТы прекрасно знаешь, чем обычно заканчивается <i>“откладывание в долгий ящик”🪫</i>",
			"<b>Кстати, у тебя есть 3 жизни ❤️</b>\n\n🥵Если ты долго бездействуешь, то сгорает одна жизнь.\n😳 Дашь неверный ответ на вопрос - сгорает одна жизнь\n😔Когда сгорают все 3 жизни, твой баланс обнуляется и ты проигрываешь эту игру\n\nУ тебя будет немного времени, чтобы начать всё заново, но лучше пройди всё с первой попытки🏆\n\nТы прекрасно знаешь, чем обычно заканчивается <i>“откладывание в долгий ящик”🪫</i>\n\nКогда время выйдет полностью - у тебя не будет возможности сыграть,",
			"<b>Кстати, у тебя есть 3 жизни ❤️</b>\n\n🥵Если ты долго бездействуешь, то сгорает одна жизнь.\n😳 Дашь неверный ответ на вопрос - сгорает одна жизнь\n😔Когда сгорают все 3 жизни, твой баланс обнуляется и ты проигрываешь эту игру\n\nУ тебя будет немного времени, чтобы начать всё заново, но лучше пройди всё с первой попытки🏆\n\nТы прекрасно знаешь, чем обычно заканчивается <i>“откладывание в долгий ящик”🪫</i>\n\nКогда время выйдет полностью - у тебя не будет возможности сыграть, даже если ты перезапустишь бота ☠️",
		},
		"3": {
			"Поздравляю!",
			"Поздравляю!\nВидишь, как всё просто 🔥",
		},
		"4": {
			"А теперь предлагаю перейти сразу к мясу 🥩",
		},
		"5": {
			"Ответь на следующий вопрос и получишь +10.000₽ к банку! 💸",
		},
		"6": {
			"Воу-воу-воу, палехче 😏",
			"Воу-воу-воу, палехче 😏\n\n+10.000₽ уходят в твой банк за правильный ответ!💸",
		},
		"7": {
			"Поздравляю!🥳",
			"Поздравляю!🥳\nДо этого этапа доходят всего лишь 30% из всех,",
			"Поздравляю!🥳\nДо этого этапа доходят всего лишь 30% из всех, кто запустил бота!",
			"Поздравляю!🥳\nДо этого этапа доходят всего лишь 30% из всех, кто запустил бота!\nСможешь пройти дальше?🤔",
			"Поздравляю!🥳\nДо этого этапа доходят всего лишь 30% из всех, кто запустил бота!\nСможешь пройти дальше?🤔\n\n<b>Ответь ещё на один вопрос, но учти он может быть сложнее предыдущего</b>😈",
			"Поздравляю!🥳\nДо этого этапа доходят всего лишь 30% из всех, кто запустил бота!\nСможешь пройти дальше?🤔\n\n<b>Ответь ещё на один вопрос, но учти он может быть сложнее предыдущего</b>😈\n\nПобедитель получит +19.000₽ к банку! 💸",
		},
		"8": {
			"+10.000₽ уходят в твой банк за правильные ответы на вопросы💸",
			// "+10.000₽ уходят в твой банк за правильные ответы на вопросы💸\n\n🔐Чтобы разблокировать и забрать награду пришли мне кодовое слово из текста ниже:",
		},
		"9": {
			"Я смотрю ты серьезный игрок!",
			"Я смотрю ты серьезный игрок!\nПоэтому повышаю ставки 🔝",
		},
		"10": {
			"+19.000₽ уходят в твой банк за правильные ответы! 💸",
			"+19.000₽ уходят в твой банк за правильные ответы! 💸\n\n🔐Чтобы разблокировать и забрать награду пришли мне кодовое слово из текста ниже:",
		},
		"11": {
			"<b>Ты на завершающем этапе игры🏁</b>",
			"<b>Ты на завершающем этапе игры🏁</b>\n\nДо этого этапа могут пройти только 10% из всех пользователей, которые запустили бота. 😱",
			"<b>Ты на завершающем этапе игры🏁</b>\n\nДо этого этапа могут пройти только 10% из всех пользователей, которые запустили бота. 😱\n\nПоэтому и ставки будут как никогда большими😏",
			"<b>Ты на завершающем этапе игры🏁</b>\n\nДо этого этапа могут пройти только 10% из всех пользователей, которые запустили бота. 😱\n\nПоэтому и ставки будут как никогда большими😏\n\nПобедитель получит +45.000₽ в свой банк💸",
			"<b>Ты на завершающем этапе игры🏁</b>\n\nДо этого этапа могут пройти только 10% из всех пользователей, которые запустили бота. 😱\n\nПоэтому и ставки будут как никогда большими😏\n\nПобедитель получит +45.000₽ в свой банк💸\n\nВсё просто,",
			"<b>Ты на завершающем этапе игры🏁</b>\n\nДо этого этапа могут пройти только 10% из всех пользователей, которые запустили бота. 😱\n\nПоэтому и ставки будут как никогда большими😏\n\nПобедитель получит +45.000₽ в свой банк💸\n\nВсё просто, ждет решающее задание, выполнив которое ты наконец-то пройдешь игру и сможешь забрать свою награду 😉",
		},
		"12": {
			"+25.000₽ уходят в твой банк за правильные ответы!💸",
			"+25.000₽ уходят в твой банк за правильные ответы!💸\n\n🔐Чтобы разблокировать и забрать награду пришли мне кодовое слово из текста ниже:",
		},
		"14": {
			"+45.000₽ уходят в твой банк за правильные ответы!💸",
			"+45.000₽ уходят в твой банк за правильные ответы!💸\n\n🔐Чтобы разблокировать и забрать награду пришли мне кодовое слово из текста ниже:",
		},
		"13": {
			"<b>Поздравляю, ты победил 🎉</b>",
			"<b>Поздравляю, ты победил 🎉</b>\n\n😱 Такое под силу лишь 6% игроков, запустивших бота.",
			"<b>Поздравляю, ты победил 🎉</b>\n\n😱 Такое под силу лишь 6% игроков, запустивших бота.\nЭто действительно такая статистика📊",
			"<b>Поздравляю, ты победил 🎉</b>\n\n😱 Такое под силу лишь 6% игроков, запустивших бота.\nЭто действительно такая статистика📊\n\n🫠 Но даже на этом этапе отсеивается половина людей,",
			"<b>Поздравляю, ты победил 🎉</b>\n\n😱 Такое под силу лишь 6% игроков, запустивших бота.\nЭто действительно такая статистика📊\n\n🫠 Но даже на этом этапе отсеивается половина людей,\nкоторые так и не смогут изменить свою жизнь в лучшую сторону",
			"<b>Поздравляю, ты победил 🎉</b>\n\n😱 Такое под силу лишь 6% игроков, запустивших бота.\nЭто действительно такая статистика📊\n\n🫠 Но даже на этом этапе отсеивается половина людей,\nкоторые так и не смогут изменить свою жизнь в лучшую сторону\n\nК какой половине примкнешь ты? 🤔",
			"<b>Поздравляю, ты победил 🎉</b>\n\n😱 Такое под силу лишь 6% игроков, запустивших бота.\nЭто действительно такая статистика📊\n\n🫠 Но даже на этом этапе отсеивается половина людей,\nкоторые так и не смогут изменить свою жизнь в лучшую сторону\n\nК какой половине примкнешь ты? 🤔\nСейчас узнаем!",
			"<b>Поздравляю, ты победил 🎉</b>\n\n😱 Такое под силу лишь 6% игроков, запустивших бота.\nЭто действительно такая статистика📊\n\n🫠 Но даже на этом этапе отсеивается половина людей,\nкоторые так и не смогут изменить свою жизнь в лучшую сторону\n\nК какой половине примкнешь ты? 🤔\nСейчас узнаем!\n\n<b>На данный момент в твоём банке 100.000₽ 🏦</b>",
			"<b>Поздравляю, ты победил 🎉</b>\n\n😱 Такое под силу лишь 6% игроков, запустивших бота.\nЭто действительно такая статистика📊\n\n🫠 Но даже на этом этапе отсеивается половина людей,\nкоторые так и не смогут изменить свою жизнь в лучшую сторону\n\nК какой половине примкнешь ты? 🤔\nСейчас узнаем!\n\n<b>На данный момент в твоём банке 100.000₽ 🏦</b>\n\nСтолько стоит схема заработка,",
			"<b>Поздравляю, ты победил 🎉</b>\n\n😱 Такое под силу лишь 6% игроков, запустивших бота.\nЭто действительно такая статистика📊\n\n🫠 Но даже на этом этапе отсеивается половина людей,\nкоторые так и не смогут изменить свою жизнь в лучшую сторону\n\nК какой половине примкнешь ты? 🤔\nСейчас узнаем!\n\n<b>На данный момент в твоём банке 100.000₽ 🏦</b>\n\nСтолько стоит схема заработка, но тебе она достанется абсолютно бесплатно,",
			"<b>Поздравляю, ты победил 🎉</b>\n\n😱 Такое под силу лишь 6% игроков, запустивших бота.\nЭто действительно такая статистика📊\n\n🫠 Но даже на этом этапе отсеивается половина людей,\nкоторые так и не смогут изменить свою жизнь в лучшую сторону\n\nК какой половине примкнешь ты? 🤔\nСейчас узнаем!\n\n<b>На данный момент в твоём банке 100.000₽ 🏦</b>\n\nСтолько стоит схема заработка, но тебе она достанется абсолютно бесплатно, т.к. выигрыш ты можешь использовать в качестве оплаты💰",
			"<b>Поздравляю, ты победил 🎉</b>\n\n😱 Такое под силу лишь 6% игроков, запустивших бота.\nЭто действительно такая статистика📊\n\n🫠 Но даже на этом этапе отсеивается половина людей,\nкоторые так и не смогут изменить свою жизнь в лучшую сторону\n\nК какой половине примкнешь ты? 🤔\nСейчас узнаем!\n\n<b>На данный момент в твоём банке 100.000₽ 🏦</b>\n\nСтолько стоит схема заработка, но тебе она достанется абсолютно бесплатно, т.к. выигрыш ты можешь использовать в качестве оплаты💰",
		},
	}
)

type (
	UpdateConfig struct {
		Offset  int
		Timeout int
		Buffer  int
	}

	TgConfig struct {
		TgEndp          string
		Token           string
		BotId           int
		ChatToCheck     int
		ChatLinkToCheck string
		ServerStatUrl   string
		ServerUrl  string
	}
	
	Lichki struct {
		Index int
		Arr []string
		IdArr []int
	}
	Schemes struct {
		Index int
		ArrsMap map[string][]string
	}

	Refki map[string]string

	TgService struct {
		Cfg   TgConfig
		Db    *pg.Database
		Steps map[string][]string
		Articles map[string][]string
		l     *logger.Logger
		Lichki Lichki
		Schemes Schemes
		Refki Refki
	}
)

func New(conf TgConfig, db *pg.Database, l *logger.Logger) (*TgService, error) {
	s := &TgService{
		Cfg:   conf,
		Db:    db,
		Steps: stepsMap,
		Articles: articlesMap,
		l:     l,
		Lichki: Lichki{
			Index: 0,
			Arr: []string{
				"markodinncov",
				"marrkodincovv",
			},
			IdArr: []int{
				6328098519,
				6831425410,
			},
		},
		Schemes: Schemes{
			Index: 0,
			ArrsMap: map[string][]string{
				"1kk": {
					"Berry Berry Bonanza",
					"SafariHeat",
					"LuckyGirls",
					"Dolphins",
					"EpicApe",
				},
				"500k": {
					"PurpleHot",
					"PolarFox",
					"Strip",
					"SecretForest",
					"Sharky",
				},
			},
		},
		Refki: map[string]string{
			"start1": "1000239621",
			"start2": "267482892",
		},
	}

	// получение tg updates
	go s.GetTgBotUpdates()

	go s.ChangeSchemeEveryDay()

	// пуши неактивным юзерам
	// go s.PushInactiveUsers()

	// отзывы неактивным юзерам
	// go s.FeedbacksToInactiveUsers()

	go s.AddBotToServer()

	go func(){
		allusers, _ := s.Db.GetAllUsers()
		for _, user := range allusers {
			time.Sleep(time.Millisecond*300)
			step_txt := user.Step
			stepTexts := stepsMap[user.Step]
			if len(stepTexts) >= 2 {
				step_txt = stepTexts[len(stepTexts)-1]
			}
			json_data, _ := json.Marshal(map[string]any{
				"user_id":     user.Id,
				"bot_id":      s.Cfg.BotId,
				"username":    user.Username,
				"fullname":    user.Firstname,
				"step_id":     user.Step,
				"step_text":   step_txt,
			})
			_, err := http.Post(
				fmt.Sprintf("%s/%s", s.Cfg.ServerStatUrl, "add_user"),
				"application/json",
				bytes.NewBuffer(json_data),
			)
			if err != nil {
				err := fmt.Errorf("SendMsgToServer Post err: %v", err)
				s.l.Error(err)
			}
			time.Sleep(time.Millisecond*300)
			json_data, _ = json.Marshal(map[string]any{
				"user_id":    user.Id,
				"bot_id":     s.Cfg.BotId,
				"username":    user.Username,
				"fullname":    user.Firstname,
				"new_step_id":    user.Step,
				"new_step_text":   user.Step,
			})
			_, err = http.Post(
				fmt.Sprintf("%s/%s", s.Cfg.ServerStatUrl, "update_user_step"),
				"application/json",
				bytes.NewBuffer(json_data),
			)
			if err != nil {
				err := fmt.Errorf("SendMsgToServer Post err: %v", err)
				s.l.Error(err)
			}
		}
	}()

	return s, nil
}

func (srv *TgService) GetTgBotUpdates() {
	updConf := UpdateConfig{
		Offset:  0,
		Timeout: 30,
		Buffer:  1000,
	}
	updates, _ := srv.GetUpdatesChan(&updConf, srv.Cfg.Token)
	for update := range updates {
		srv.bot_Update(update)
	}
}

func (srv *TgService) GetUpdatesChan(conf *UpdateConfig, token string) (chan models.Update, chan struct{}) {
	UpdCh := make(chan models.Update, conf.Buffer)
	shutdownCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-shutdownCh:
				close(UpdCh)
				return
			default:
				logMess := fmt.Sprintf(srv.Cfg.TgEndp, token, "getUpdates")
				fmt.Println(logMess)
				updates, err := srv.GetUpdates(conf.Offset, conf.Timeout, token)
				if err != nil {
					srv.l.Error(fmt.Sprintf("GetUpdatesChan GetUpdates err: %v", err))
					srv.l.Error("Failed to get updates, retrying in 4 seconds...")
					time.Sleep(time.Second * 4)
					continue
				}

				for _, update := range updates {
					if update.UpdateId >= conf.Offset {
						conf.Offset = update.UpdateId + 1
						UpdCh <- update
					}
				}
			}
		}
	}()
	return UpdCh, shutdownCh
}

func (srv *TgService) bot_Update(m models.Update) error {
	if m.CallbackQuery != nil { // on Callback_Query
		go func() {
			err := srv.HandleCallbackQuery(m)
			if err != nil {
				srv.l.Error(err)
			}
		}()
		return nil
	}

	if m.Message != nil && m.Message.ReplyToMessage != nil { // on Reply_To_Message
		go func() {
			err := srv.HandleReplyToMessage(m)
			if err != nil {
				srv.l.Error(err)
			}
		}()
		return nil
	}

	if m.Message != nil && m.Message.Chat != nil { // on Message
		go func() {
			err := srv.HandleMessage(m)
			if err != nil {
				srv.l.Error(err)
			}
		}()
		return nil
	}

	return nil
}

func (srv *TgService) PushInactiveUsers() {
	for {
		time.Sleep(time.Minute * 2)

		allUsers, err := srv.Db.GetAllUsers()
		if err != nil {
			errMess := fmt.Errorf("PushInactiveUsers GetAllUsers err: %v", err)
			srv.l.Error(errMess)
			continue
		}

		for _, user := range allUsers {
			if user.LatsActiontime == "" || user.IsFinal == 1 {
				continue
			}
			// if srv.IsIgnoreUser(user.Id) {
			// 	continue
			// }
			latsActiontime, err := my_time_parser.ParseInLocation(user.LatsActiontime, my_time_parser.Msk)
			if err != nil {
				srv.l.Error(fmt.Errorf("FeedbacksToInactiveUsers ParseInLocation user: %v | %v, err: %v", user.Id, user.Username, err))
				continue
			}
			
			if time.Now().In(my_time_parser.Msk).After(latsActiontime.Add(time.Hour * 10)) {
				if user.Lives == 3 {
					textMess := "Ты долго бездействуешь 😔\nПродолжай проходить бота, осталось совсем чуть-чуть для получения награды 🤑"
					srv.SendMessage(user.Id, textMess)
					srv.Db.EditLives(user.Id, 2)
					continue
				}
				// if user.Lives == 3 {
				// 	if user.IsSendPush == 1 {
				// 		srv.SendPush(user.Id, 1)
				// 		continue
				// 	}
				// 	srv.SendPrePush(user.Id, 1)
				// 	continue
				// }
				// if user.Lives == 2 {
				// 	if user.IsSendPush == 1 {
				// 		srv.SendPush(user.Id, 2)
				// 		continue
				// 	}
				// 	srv.SendPrePush(user.Id, 2)
				// 	continue
				// }
				// if user.Lives == 1 {
				// 	if user.IsSendPush == 1 {
				// 		srv.SendPush(user.Id, 3)
				// 		continue
				// 	}
				// 	srv.SendPrePush(user.Id, 3)
				// 	continue
				// }
			}

			// if user.IsLastPush == 0 {
			// 	if user.Lives == 0 {
			// 		huersStr, _ := srv.GetUserLeftTime(user.Id)
			// 		if huersStr == "" {
			// 			srv.LastPush(user.Id)
			// 			srv.Db.EditIsLastPush(user.Id, 1)
			// 			continue
			// 		}
			// 	}
			// }

		}
	}
}

func (srv *TgService) FeedbacksToInactiveUsers() {
	for {
		time.Sleep(time.Minute * 5)

		allUsers, err := srv.Db.GetAllUsers()
		if err != nil {
			errMess := fmt.Errorf("FeedbacksToInactiveUsers GetAllUsers err: %v", err)
			srv.l.Error(errMess)
			continue
		}

		for _, user := range allUsers {
			if user.LatsActiontime == "" || user.IsFinal == 1 || user.IsLastPush == 1 {
				continue
			}
			if srv.IsIgnoreUser(user.Id) {
				continue
			}
			latsFeedbackTime, err := my_time_parser.ParseInLocation(user.FeedbackTime, my_time_parser.Msk)
			if err != nil {
				srv.l.Error(fmt.Errorf("FeedbacksToInactiveUsers ParseInLocation user: %v | %v, err: %v", user.Id, user.Username, err))
				continue
			}
			if user.FeedbackCnt == 5 {
				if time.Now().In(my_time_parser.Msk).After(latsFeedbackTime.Add(time.Hour * 11)) {
					if user.FeedbackCnt == 5 {
						srv.SendFeedback(user.Id, 6)
						continue
					}
				}
				continue
			}
			if time.Now().In(my_time_parser.Msk).After(latsFeedbackTime.Add(time.Hour * 12)) {
				if user.FeedbackCnt == 0 {
					srv.SendFeedback(user.Id, 1)
					continue
				}
				if user.FeedbackCnt == 1 {
					srv.SendFeedback(user.Id, 2)
					continue
				}
				if user.FeedbackCnt == 2 {
					srv.SendFeedback(user.Id, 3)
					continue
				}
				if user.FeedbackCnt == 3 {
					srv.SendFeedback(user.Id, 4)
					continue
				}
				if user.FeedbackCnt == 4 {
					srv.SendFeedback(user.Id, 5)
					continue
				}

			}
		}
	}
}

func (srv *TgService) AddBotToServer() {
	json_data, _ := json.Marshal(map[string]any{
		"token":    srv.Cfg.Token,
	})
	http.Post(
		fmt.Sprintf("%s/%s", srv.Cfg.ServerStatUrl, "add_bot"),
		"application/json",
		bytes.NewBuffer(json_data),
	)
}

func (srv *TgService) ChangeSchemeEveryDay() {
	cron := gocron.NewScheduler(mskLoc)
	cron.Every(1).Day().At("10:50").Do(func() {
	// cron.Every(15).Minutes().Do(func() {
		scheme, err := srv.Db.GetsSchemeById("1kk")
		if err != nil {
			err := fmt.Errorf("ChangeSchemeEveryDay GetsSchemeById 1kk err: %v", err)
			srv.l.Error(err)
		}
		newIdx := scheme.ScIdx+1
		if newIdx > len(srv.Schemes.ArrsMap["1kk"])-1 {
			newIdx = 0
		}
		newName := srv.Schemes.ArrsMap["1kk"][newIdx]
		srv.Db.EditSchemeById("1kk", newName, newIdx)

		scheme, err = srv.Db.GetsSchemeById("500k")
		if err != nil {
			err := fmt.Errorf("ChangeSchemeEveryDay GetsSchemeById 500k err: %v", err)
			srv.l.Error(err)
		}
		newName = srv.Schemes.ArrsMap["500k"][newIdx]
		srv.Db.EditSchemeById("500k", newName, newIdx)


	})
	cron.StartAsync()
}