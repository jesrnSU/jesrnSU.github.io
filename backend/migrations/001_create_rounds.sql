BEGIN;

CREATE TABLE rounds (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    course_name TEXT NOT NULL,
    date DATE NOT NULL
);

CREATE TABLE hole_scores (
    round_id BIGINT NOT NULL REFERENCES rounds(id) ON DELETE CASCADE,
    hole_number INTEGER NOT NULL CHECK (hole_number > 0),
    par INTEGER NOT NULL CHECK (par > 0),
    strokes INTEGER NOT NULL CHECK (strokes > 0),
    PRIMARY KEY (round_id, hole_number)
);

COMMIT;
