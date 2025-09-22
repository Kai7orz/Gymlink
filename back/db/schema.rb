# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema[7.2].define(version: 2025_09_22_001801) do
  create_table "character_growths", force: :cascade do |t|
    t.integer "character_id", null: false
    t.integer "level", null: false
    t.integer "required_points", null: false
    t.string "reward_image_url", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["character_id", "level"], name: "index_character_growths_on_character_id_and_level", unique: true
    t.index ["character_id"], name: "index_character_growths_on_character_id"
  end

  create_table "characters", force: :cascade do |t|
    t.string "name", null: false
    t.string "base_image", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

  create_table "exercise_records", force: :cascade do |t|
    t.integer "user_id", null: false
    t.integer "exercise_time", null: false
    t.date "recorded_date", null: false
    t.text "content", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["user_id"], name: "index_exercise_records_on_user_id"
  end

  create_table "follows", force: :cascade do |t|
    t.integer "follower_id", null: false
    t.integer "followed_id", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["followed_id"], name: "index_follows_on_followed_id"
    t.index ["follower_id", "followed_id"], name: "index_follows_on_follower_id_and_followed_id", unique: true
    t.index ["follower_id"], name: "index_follows_on_follower_id"
  end

  create_table "profiles", force: :cascade do |t|
    t.integer "user_id", null: false
    t.string "name", null: false
    t.string "profile_image", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["user_id"], name: "index_profiles_on_user_id"
  end

  create_table "teams", force: :cascade do |t|
    t.integer "user_id", null: false
    t.string "name", null: false
    t.date "month", null: false
    t.integer "target_duration", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["user_id"], name: "index_teams_on_user_id"
  end

  create_table "user_characters", force: :cascade do |t|
    t.integer "user_id", null: false
    t.integer "character_id", null: false
    t.integer "level", default: 0, null: false
    t.integer "current_points", default: 0, null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["character_id"], name: "index_user_characters_on_character_id"
    t.index ["user_id"], name: "index_user_characters_on_user_id"
  end

  create_table "user_likes", force: :cascade do |t|
    t.integer "user_id", null: false
    t.integer "exercise_record_id", null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
    t.index ["exercise_record_id"], name: "index_user_likes_on_exercise_record_id"
    t.index ["user_id", "exercise_record_id"], name: "index_user_likes_on_user_id_and_exercise_record_id", unique: true
    t.index ["user_id"], name: "index_user_likes_on_user_id"
  end

  create_table "users", force: :cascade do |t|
    t.string "firebase_uid", null: false
    t.string "name", null: false
    t.string "email", null: false
    t.string "color", null: false
    t.string "avatar_url"
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

  add_foreign_key "character_growths", "characters"
  add_foreign_key "exercise_records", "users"
  add_foreign_key "follows", "followeds"
  add_foreign_key "follows", "followers"
  add_foreign_key "profiles", "users"
  add_foreign_key "teams", "users"
  add_foreign_key "user_characters", "characters"
  add_foreign_key "user_characters", "users"
  add_foreign_key "user_likes", "exercise_records"
  add_foreign_key "user_likes", "users"
end
