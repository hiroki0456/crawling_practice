create TABLE Users (
    id int auto_increment primary key not null,
    user_id char(255) unique not null,
    crawling_date timestamp not null
);

create TABLE Banks(
    id int auto_increment primary key not null,
    name char(255) not null,
    updated_at timestamp not null,
    user_id int not null,
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE TABLE Details (
    id int auto_increment primary key not null,
    bankName char(255) not null,
    tradingContent char(255),
    amount int not null,
    Updated_at timestamp not null,
    user_id int not null,
    bank_id int not null,
    FOREIGN KEY (user_id) REFERENCES Users (id),
    FOREIGN KEY (bank_id) REFERENCES Banks (id)
);