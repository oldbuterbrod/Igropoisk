CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE games (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    developer TEXT,
    platforms TEXT,
    publisher TEXT,
    release_date DATE,
    cover_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE game_genres (
    game_id INTEGER REFERENCES games(id) ON DELETE CASCADE,
    genre_id INTEGER REFERENCES genres(id) ON DELETE CASCADE,
    PRIMARY KEY (game_id, genre_id)
);


CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    game_id INTEGER REFERENCES games(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CHECK (length(text) >= 10 AND length(text) <= 189)
);


CREATE TABLE ratings (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    game_id INTEGER REFERENCES games(id) ON DELETE CASCADE,
    score INTEGER CHECK (score >= 1 AND score <= 10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, game_id)
);


CREATE TABLE favorites (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    game_id INTEGER REFERENCES games(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, game_id)
);

CREATE UNIQUE INDEX unique_review_per_user_per_game
ON reviews(user_id, game_id);



INSERT INTO games (
    id, title, description, developer, platforms, publisher, release_date, cover_url, created_at
) VALUES
(1, 'Resident Evil 2 Remake',
'Ремейк классического survival horror, где игроки управляют Леоном Кеннеди и Клэр Редфилд, оказавшимися в Раккун-Сити во время зомби-апокалипсиса. Сюжет разделён на кампании для каждого персонажа, предлагая уникальные перспективы и встречи. Геймплей сочетает стрельбу, решение головоломок и управление ресурсами, с обновлённой камерой от третьего лица и улучшенной графикой.',
'Capcom', 'PC, PS4, Xbox One', 'Capcom', '2019-01-25', 
'https://4.downloader.disk.yandex.ru/preview/2cdfaf58fbe8307f6dc68cef7ca302888651319da4cbad5c0166e03395653109/inf/bH-xL3soHOqkDWDZBiaav_bFadY5vfZxrhDaDqPO-6YL6dYboZJWpRwveRqfVwvB5GIjLKJt5JUejFcpPNd8mA%3D%3D?uid=208664216&filename=Resident_Evil_2_Remake.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=1920x922', CURRENT_TIMESTAMP),

(2, 'Silent Hill 2 Remake',
'Психологический хоррор, в котором Джеймс Сандерленд получает письмо от покойной жены и отправляется в туманный город Сайлент Хилл. Игра исследует темы вины и скорби, предлагая мрачную атмосферу, сложные головоломки и встречи с монстрами, олицетворяющими внутренние демоны героя.',
'Bloober Team', 'PC, PS5, Xbox Series X', 'Konami', '2024-10-10', 
'https://3.downloader.disk.yandex.ru/preview/9217a078903c6e4d5020fe7daf5e088db7d25620788936c95759ae2823abefb9/inf/Nrj8zy3HHYL_fwf2zbK8UTJix91Z9mEJXdRDJhOqM4dAOYIh0b0oY9UgzN7ijoAjsTa9WIGjMHxBIY4xqVZ3AA%3D%3D?uid=208664216&filename=silent-hill-2-remake_87rs.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=1920x922', CURRENT_TIMESTAMP),

(3, 'Alan Wake',
'Психологический триллер, рассказывающий о писателе Алане Уэйке, который ищет свою пропавшую жену в городе Брайт-Фоллс. Игра разделена на эпизоды, как телесериал, и сочетает исследование, использование света как оружия против тьмы и элементы хоррора.',
'Remedy Entertainment', 'PC, Xbox 360', 'Microsoft Game Studios', '2010-05-14', 
'https://2.downloader.disk.yandex.ru/preview/27a6db99fe096ca8548bbe164e92ec9dec0625bd982892480455ff80a1eba835/inf/cSoktQ8kbiWfVucJ86eP_Wpo-EoDwq1KF2zlFyLCP9OG1pjK6pPXigcw0QHh51IC1nkOl3lEnLL6PDvkoiR76w%3D%3D?uid=208664216&filename=images.jfif&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=1920x922', CURRENT_TIMESTAMP),

(4, 'The Last of Us',
'Постапокалиптическое приключение, где контрабандист Джоэл сопровождает подростка Элли через заражённые США. Игра сочетает стелс, боевые элементы и эмоциональный сюжет, исследуя темы выживания, доверия и жертвы.',
'Naughty Dog', 'PS3, PS4, PS5, PC', 'Sony Computer Entertainment', '2013-06-14',
 'https://2.downloader.disk.yandex.ru/preview/957099e1b6919386a27174888ae281985a43855ff1c6b32f04beb38f91cd1927/inf/nlIVE1pF-oN24pfThHrryMAI7UdHxUDOwyo-7yyv9r6g2MgLqOlroFifXqmdFCMqHXp4zt68FrkMHsTFWzSJrg%3D%3D?uid=208664216&filename=EGS_TheLastofUsPartI_NaughtyDogLLC_S2_1200x1600-41d1b88814bea2ee8cb7986ec24713e0.jfif&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=1920x922', CURRENT_TIMESTAMP),

INSERT INTO games (id, title, description, developer, platforms, publisher, release_date, cover_url, created_at) VALUES
(5, 'Red Dead Redemption 2', 
'Эпический приключенческий боевик от Rockstar Games, действие которого разворачивается на Диком Западе в 1899 году. Игроки управляют Артуром Морганом, членом банды Ван дер Линде, которая пытается выжить в условиях наступления цивилизации. Игра сочетает открытый мир, реалистичную механику выживания, моральные выборы и захватывающий сюжет.', 
'Rockstar Games', 'PC, PS4, Xbox One', 'Rockstar Games', '2018-10-26', 
'https://1.downloader.disk.yandex.ru/preview/37902bb7d0a25d67b891ad3f90a737ee302b48472cae7857313fc5b65b9ac703/inf/ip9wKzTraj-HfGhuQyaIamvsQqqJtL3KvbTxjQBPkJitQ6tet7nEJygaJ0Nq1bFloAwtbwAwgXQZ_HOl-2rMiQ%3D%3D?uid=208664216&filename=Red_Dead_Redemption_2_coverart.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=1901x922', 
CURRENT_TIMESTAMP),

(6, 'Grand Theft Auto V', 
'Легендарная криминальная сага, переносящая игроков в вымышленный город Лос-Сантос. Сюжет вращается вокруг трёх преступников — Майкла, Франклина и Тревора, чьи судьбы переплетаются во время серии опасных ограблений. Помимо насыщенного сюжета, игра предлагает огромный открытый мир и многопользовательский режим GTA Online.', 
'Rockstar North', 'PC, PS4, PS5, Xbox One, Xbox Series X', 'Rockstar Games', '2013-09-17', 
'https://3.downloader.disk.yandex.ru/preview/74d237e6af372af71639ea54748821d87be90f500650a258ef4872d58680cde3/inf/xDTDV7QOyCaHASfiEYUEEN1O07XZp45HgrSlTBdbCqYhR_GcvERmY3EaHuEpT0_c9p1g5C8LGAtyZ_oBDcSoGA%3D%3D?uid=208664216&filename=GTAV_Official_Cover_Art.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=940x902', 
CURRENT_TIMESTAMP),

(7, 'Cyberpunk 2077', 
'Футуристическая RPG от создателей The Witcher 3, действие которой происходит в мрачном мегаполисе Найт-Сити. Игрок берёт на себя роль Ви — наёмника, который ввязывается в опасную авантюру в поисках чипа, дарующего бессмертие. Игра предлагает нелинейный сюжет, кибернетические улучшения и множество способов прохождения.', 
'CD Projekt Red', 'PC, PS5, Xbox Series X', 'CD Projekt', '2020-12-10', 
'https://3.downloader.disk.yandex.ru/preview/21ada874f8ea4324d32f1fdb532e8f58203ce1a79a144f43a3697378877df303/inf/2ue7DE89lKqapJutVBUG3BGq7PkVzQZGAymR8zUiJ8e4VeKkVkrjkuRCac4kyC51s5qXs2uqnl59tn2nbl9Rmg%3D%3D?uid=208664216&filename=%D0%9E%D0%B1%D0%BB%D0%BE%D0%B6%D0%BA%D0%B0_%D0%BA%D0%BE%D0%BC%D0%BF%D1%8C%D1%8E%D1%82%D0%B5%D1%80%D0%BD%D0%BE%D0%B9_%D0%B8%D0%B3%D1%80%D1%8B_Cyberpunk_2077.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=940x902', 
CURRENT_TIMESTAMP),

(8, 'The Witcher 3: Wild Hunt', 
'Одна из лучших RPG всех времён, где игроки становятся Геральтом из Ривии — ведьмаком, охотящимся на чудовищ. В поисках своей приёмной дочери Цири Геральт путешествует по огромному миру, сталкиваясь с политическими интригами, моральными дилеммами и опасными существами. Игра славится глубокими персонажами, богатым сюжетом и атмосферными локациями.', 
'CD Projekt Red', 'PC, PS4, Xbox One, Switch', 'CD Projekt', '2015-05-19', 
'https://2.downloader.disk.yandex.ru/preview/2b5e46c310e25d70dda4b7cbe653e46779b4a93e339bbad00688fa5f62663cd4/inf/rAsAOF_tKjnUMZHlxj7Cw0GKh3etAFFoDhTypMfoTrRAuFN9oG3jUXOQUou69DeNQ-pbHF5rvGCptlKuiKmBaA%3D%3D?uid=208664216&filename=The_Witcher_3-_Wild_Hunt_Cover.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=940x902', 
CURRENT_TIMESTAMP),

(9, 'Fallout 4',
'Постапокалиптическая RPG от Bethesda, действие которой происходит в разрушенном ядерной войной мире. Игрок управляет выжившим из Убежища 111, исследуя пустоши, строя поселения и принимая судьбоносные решения.',
'Bethesda Game Studios', 'PC, PS4, Xbox One', 'Bethesda Softworks', '2015-11-10',
'https://3.downloader.disk.yandex.ru/preview/0e998843d26c484b6625c721375059162d85456c9d247713bdc8279b65e36d3f/inf/X4AWlr0MZUddgPbxozOMDii5fypf4mup6FtI3N_nZnD3SxRhtdTzfuWxPiui35kGnB8QuWKHQTAhuaO1prXocQ%3D%3D?uid=208664216&filename=81LVtnitzDL.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=1901x922',
CURRENT_TIMESTAMP),

(10, 'Mass Effect: Legendary Edition', 
'Ремастер культовой космической трилогии, где игроки принимают на себя роль капитана Шепарда. В борьбе за выживание галактики перед угрозой Жнецов игроку предстоит принимать ключевые решения, которые влияют на сюжет и судьбы персонажей. В Legendary Edition улучшена графика и объединены все три части с DLC.', 
'BioWare', 'PC, PS4, Xbox One', 'Electronic Arts', '2021-05-14', 
'https://3.downloader.disk.yandex.ru/preview/b83fe5bce6e7f893be360fe22562927bc2fac50ba0cf8df9b34e6f2182635c43/inf/PLG7a0yI9pMRQzcHrP63iczM8AbLnImv5fxfLaBE0XGvlYnc0ISRHJujz5RrW5l0VkiSQkB4lD-MSUGzzvasKA%3D%3D?uid=208664216&filename=%D0%9E%D0%B1%D0%BB%D0%BE%D0%B6%D0%BA%D0%B0_Mass_Effect_Legendary_Edition.png&disposition=inline&hash=&limit=0&content_type=image%2Fpng&owner_uid=208664216&tknv=v3&size=940x902', 
CURRENT_TIMESTAMP),

(11, 'Doom: The Dark Age', 
'Новая часть легендарного шутера, переносящая действие в мрачное средневековье. Игроки вновь возьмут на себя роль Думгая, сражаясь с полчищами демонов при помощи арсенала футуристического и средневекового оружия. Ожидается агрессивный геймплей и кинематографичная подача.', 
'id Software', 'PC, Xbox Series X, PS5', 'Bethesda', '2025-05-15', 
'https://4.downloader.disk.yandex.ru/preview/9a409463775589f0d5409afd83ee1f5ab6335d3297c7f521b82f7ab031e8a738/inf/fR_gS2WKDDU0E-B4Udt6IwQ7XNi3YksnOa516Z0KdOwXqxEKtBZWOi5mC1Aepf1KkeP7C9RboKQGZaswIAocrQ%3D%3D?uid=208664216&filename=DOOM_-_The_Dark_Ages.webp.png&disposition=inline&hash=&limit=0&content_type=image%2Fpng&owner_uid=208664216&tknv=v3&size=940x902', 
CURRENT_TIMESTAMP),

(12, 'Warhammer 40,000: Space Marine 2', 
'Экшен от третьего лица, где игроки становятся космодесантником, сражающимся с орками и тиранидами. Игра делает упор на динамичные бои, мощные атаки и эпические сражения в сеттинге вселенной Warhammer 40K.', 
'Saber Interactive', 'PC, PS5, Xbox Series X', 'Focus Entertainment', '2024-09-09', 
'https://3.downloader.disk.yandex.ru/preview/fa285bbfd443d20b295d18d389822e23d2af0ccd6c09f90cda7782afe87865dc/inf/TWnmU5R4BrtgFGBcsg5A2rHlMBTb2lPHdrUO8NSlYADr6-J6gkvd09JqX5nil4oRSyXPiP73rUZf205EwtNiLw%3D%3D?uid=208664216&filename=Space_Marine_II_cover_art.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=940x902', 
CURRENT_TIMESTAMP),

(13, 'Mafia: The Old Country', 
'Приквел к серии Mafia, рассказывающий о становлении мафиозных кланов в Америке 1920-х годов. Игра сочетает криминальный сюжет, открытый мир и тактические перестрелки.', 
'Hangar 13', 'PC, PS5, Xbox Series X', '2K Games', '2025-08-08', 
'https://3.downloader.disk.yandex.ru/preview/5a0261936be9a569b378de7b48b6062fc5a3dc7c5554e0b1479da9bd447c2104/inf/gUrcq5z5YcSo12BbrdBbbLNvyfw59XoQDN3wW-uGvDcmCJAmdTX0mD2hnp9Z_vrLpgq4HhxXE-nIXjf4jQVr3Q%3D%3D?uid=208664216&filename=Mafia_-_The_Old_Country.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=940x902', 
CURRENT_TIMESTAMP),

(14, 'Grand Theft Auto VI', 
'Долгожданное продолжение серии GTA, действие которого разворачивается в вымышленном городе Вайс-Сити. Сюжет повествует о дуэте преступников Лусии Каминос и Джейсона Дюваля в стиле истории о банде Бонни и Клайда.', 
'Rockstar North', 'PS5, Xbox Series X, PC', 'Rockstar Games', '2025-05-26', 
'https://3.downloader.disk.yandex.ru/preview/ca7c44dc7f93223b38f499c9e5f17bdf2addd3d46ddd60208789dd5ad9ff1b92/inf/i-BaICSTo_b-SszVX7-lxvKZ7RtqfXfMyQPjQd86ldwVV-Zi8hthfhT49xrUtInYYDjiPsxTLZcONjLSiJi5Ug%3D%3D?uid=208664216&filename=GTA_VI_Poster.jpg&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=940x902', 
CURRENT_TIMESTAMP),

(15, 'Resident Evil: Requiem', 
'Финальная глава глава в серии Resident Evil. Расследуя загадочные смерти людей от некой болезни, Грейс отправляется в Раккун-Сити, а именно в отель Вренвуд — туда, где была убита ее мать. Впоследствии Грейс сама становится жертвой неких мутантов.', 
'Capcom', 'PC, PS5, Xbox Series X', 'Capcom', '2026-02-27', 
'https://3.downloader.disk.yandex.ru/preview/979b4b4e77f97c517fb8489a4369a8b5a7450d6cb635cf7d1dc76969d5a82e9c/inf/IL8qVO-UEDkjdb7ePzPV6-lYy6JDcbJOtV8mCY3UkJtOXVanViNKPFN9hcZSEn4bMaMWwm-pclsSEjJ9NfqTtQ%3D%3D?uid=208664216&filename=%D0%9E%D0%B1%D0%BB%D0%BE%D0%B6%D0%BA%D0%B0_Resident_Evil_Requiem.webp&disposition=inline&hash=&limit=0&content_type=image%2Fjpeg&owner_uid=208664216&tknv=v3&size=940x902', 
CURRENT_TIMESTAMP);

INSERT INTO genres (id,name) VALUES (1,'Хоррор');
INSERT INTO game_genres (game_id, genre_id) VALUES
(1, 1),
(2, 1),
(3, 1),
(4, 1);
