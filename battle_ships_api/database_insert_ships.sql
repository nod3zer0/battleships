LET SHIPS = [ { "name": "Carrier", "placed": "false", "position": [0, 0], "parts": [ [0, 0], [0, 1], [0, 2], [0, 3], [0, 4] ] }, { "name": "Battleship", "placed": "false", "position": [0, 0], "parts": [ [0, 0], [0, 1], [0, 2], [0, 3] ] }, { "name": "Cruiser", "placed": "false", "position": [0, 0], "parts": [ [0, 0], [0, 1], [0, 2] ] }, { "name": "Submarine", "placed": "false", "position": [0, 0], "parts": [ [0, 0], [0, 1], [0, 2] ] }, { "name": "Destroyer", "placed": "false", "position": [0, 0], "parts": [ [0, 0], [0, 1] ] } ];

//sql code to insert data from that pseudo code

INSERT INTO SHIP (
    NAME
) VALUES (
    'Carrier'
);

INSERT INTO SHIP (
    NAME
) VALUES (
    'Battleship'
);

INSERT INTO SHIP (
    NAME
) VALUES (
    'Cruiser'
);

INSERT INTO SHIP (
    NAME
) VALUES (
    'Submarine'
);

INSERT INTO SHIP (
    NAME
) VALUES (
    'Destroyer'
);

INSERT INTO SHIP_PART (
    SHIP_ID,
    POS_X,
    POS_Y
) VALUES (
    1,
    0,
    0
),
(
    1,
    0,
    1
),
(
    1,
    0,
    2
),
(
    1,
    0,
    3
),
(
    1,
    0,
    4
),
(
    2,
    0,
    0
),
(
    2,
    0,
    1
),
(
    2,
    0,
    2
),
(
    3,
    0,
    3
),
(
    3,
    0,
    0
),
(
    3,
    0,
    1
),
(
    3,
    0,
    2
),
(
    4,
    0,
    0
),
(
    4,
    0,
    1
),
(
    4,
    0,
    2
),
(
    5,
    0,
    0
),
(
    5,
    0,
    1
);