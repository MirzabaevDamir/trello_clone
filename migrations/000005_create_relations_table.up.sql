ALTER TABLE
    "cards" ADD CONSTRAINT "cards_board_id_foreign" FOREIGN KEY("board_id") REFERENCES "boards"("id");
ALTER TABLE
    "boards" ADD CONSTRAINT "boards_work_space_id_foreign" FOREIGN KEY("work_space_id") REFERENCES "work_spaces"("id");
ALTER TABLE
    "work_spaces" ADD CONSTRAINT "work_spaces_owner_id_foreign" FOREIGN KEY("owner_id") REFERENCES "users"("id");

