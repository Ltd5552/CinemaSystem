CREATE TABLE IF NOT EXISTS User
(
    uid VARCHAR(20) PRIMARY KEY ,
    name VARCHAR(20) NOT NULL ,
    sex VARCHAR(5) ,
    birthday VARCHAR(20),
    location VARCHAR(20),
    phone VARCHAR(11),
    password VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS Cinema
(
    cinemaNum VARCHAR(20) PRIMARY KEY ,
    cinemaName VARCHAR(15) NOT NULL ,
    city VARCHAR(15) NOT NULL ,
    contact VARCHAR(11) NOT NULL ,
    aveCinemaScore FLOAT
);


CREATE TABLE IF NOT EXISTS Theater
(
    theaterNum VARCHAR(20) PRIMARY KEY ,
    cinemaNum VARCHAR(20),
    seatsNum INT NOT NULL ,
    FOREIGN KEY (cinemaNum) REFERENCES Cinema(cinemaNum)
);

CREATE TABLE IF NOT EXISTS Movie
(
    movieNum VARCHAR(20) PRIMARY KEY ,
    movieTitle VARCHAR(15) NOT NULL ,
    releaseDate VARCHAR(20),
    duration FLOAT,
    aveFilmScore FLOAT
);


CREATE TABLE IF NOT EXISTS Screenings
(
    screeningNum VARCHAR(20) PRIMARY KEY ,
    movieNum VARCHAR(20),
    theaterNum VARCHAR(20),
    showTime VARCHAR(20) NOT NULL ,
    remainSeats INT NOT NULL ,
    FOREIGN KEY (movieNum) REFERENCES Movie(movieNum),
    FOREIGN KEY (theaterNum) REFERENCES Theater(theaterNum)
);


CREATE TABLE IF NOT EXISTS Ticket
(
    ticketNum VARCHAR(20) PRIMARY KEY ,
    price FLOAT NOT NULL ,
    screeningNum VARCHAR(20) NOT NULL ,
    FOREIGN KEY (screeningNum) REFERENCES Screenings(screeningNum)
);


CREATE TABLE IF NOT EXISTS Buy
(
    uid VARCHAR(20),
    ticketNum VARCHAR(20),
    buyTime VARCHAR(20) NOT NULL ,
    PRIMARY KEY (uid,ticketNum),
    FOREIGN KEY (uid) REFERENCES User(uid),
    FOREIGN KEY (ticketNum) REFERENCES Ticket(ticketNum)
);


CREATE TABLE IF NOT EXISTS TicketMachine
(
    ticketMachineNum VARCHAR(20) PRIMARY KEY ,
    remainingTickets INT
);


CREATE TABLE IF NOT EXISTS Print
(
    ticketMachineNum VARCHAR(20),
    ticketNum VARCHAR(20),
    printTime VARCHAR(20) NOT NULL ,
    PRIMARY KEY (ticketNum,ticketMachineNum),
    FOREIGN KEY (ticketMachineNum) REFERENCES TicketMachine(ticketMachineNum),
    FOREIGN KEY (ticketNum) REFERENCES Ticket(ticketNum)
);

CREATE TABLE IF NOT EXISTS Evaluation
(
    evaluationId VARCHAR(20) PRIMARY KEY ,
    cinemaScore INT NOT NULL ,
    fileScore INT NOT NULL,
    movieNum VARCHAR(20) NOT NULL ,
    cinemaNum VARCHAR(20) NOT NULL ,
    FOREIGN KEY (cinemaNum) REFERENCES cinema(cinemaNum),
    FOREIGN KEY (movieNum) REFERENCES movie(movieNum)
);

CREATE TABLE IF NOT EXISTS Releases
(
    uid VARCHAR(20),
    evaluationId VARCHAR(20),
    releaseTime VARCHAR(20) NOT NULL ,
    PRIMARY KEY (uid,evaluationId),
    FOREIGN KEY (uid) REFERENCES User(uid),
    FOREIGN KEY (evaluationId) REFERENCES Evaluation(evaluationId)
);




