-- +goose Up
-- +goose StatementBegin
CREATE TABLE puzzles
(
    word     VARCHAR(15) NOT NULL PRIMARY KEY,
    category VARCHAR(60) NOT NULL,
    hint     VARCHAR(265) NULL
);

INSERT INTO puzzles ("word", "category", "hint") VALUES
('пират','криминал','Представители этой профессии занимаются незаконной деятельностью на морских судах'),
('аббат','религия','Настоятель мужского католического монастыря'),
('борец','спорт','Спортсмен учавствующий в рукопашной схватке, в которой каждый старается осилить другого, свалив его с ног'),
('вилка','кухня','Столовый прибор'),
('дебет','финансы','Левая часть счета бухгалтерского учета для отражения хозяйственной операции методом двойной записи'),
('залив','природа','часть океана, моря, озера или другого водоёма, глубоко вдающаяся в сушу, но имеющая свободный водообмен с основной частью водоёма'),
('камин','дом и уют','разновидность печи'),
('комар','насекомые','Два крыла, любит кровь');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE puzzles;
-- +goose StatementEnd