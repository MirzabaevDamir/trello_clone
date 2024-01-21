ALTER TABLE "cards" DROP CONSTRAINT "cards_board_id_foreign";
ALTER TABLE "cards" DROP CONSTRAINT "cards_user_ids_foreign";
ALTER TABLE "boards" DROP CONSTRAINT "boards_work_space_id_foreign";
ALTER TABLE "work_spaces" DROP CONSTRAINT "work_spaces_owner_id_foreign";
