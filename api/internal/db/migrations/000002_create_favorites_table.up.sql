CREATE TABLE favorites(
                          id SERIAL PRIMARY KEY,

                          session_id VARCHAR(255),

                          user_id BIGINT,

                          book_id BIGINT NOT NULL,

                          created_at TIMESTAMP DEFAULT NOW(),

                          CONSTRAINT favorites_book_fk
                              FOREIGN KEY(book_id)
                                  REFERENCES books(id)
                                  ON DELETE CASCADE,

                          CONSTRAINT favorites_user_fk
                              FOREIGN KEY(user_id)
                                  REFERENCES users(id)
                                  ON DELETE CASCADE,

                          CONSTRAINT favorites_unique_user
                              UNIQUE(user_id, book_id),

                          CONSTRAINT favorites_unique_session
                              UNIQUE(session_id, book_id)
);