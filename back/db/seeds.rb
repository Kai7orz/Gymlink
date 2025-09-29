# This file should ensure the existence of records required to run the application in every environment (production,
# development, test). The code here should be idempotent so that it can be executed at any point in every environment.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).
#
# Example:
#
#   ["Action", "Comedy", "Drama", "Horror"].each do |genre_name|
#     MovieGenre.find_or_create_by!(name: genre_name)
#   end

require "faker"
require "date"
require "securerandom"

dummy_team = Team.create!(
  name: "Dummy Team",
  target_duration: 0,
  month: Date.today
)

User.find_or_create_by!(
  firebase_uid: "uid_001",
  email: "taro@example.com"
) do |user|
  user.name = "Taro Yamada"
  user.color = "#ff0000"
  user.avatar_url = "https://example.com/avatar1.png"
  user.team_id = dummy_team.id
  user.password = "password123"
end

User.find_or_create_by!(
  firebase_uid: "uid_002",
  email: "hanako@example.com"
) do |user|
  user.name = "Hanako Suzuki"
  user.color = "#00ff00"
  user.avatar_url = "https://example.com/avatar2.png"
  user.team_id = dummy_team.id
  user.password = "password1234"
end



puts "Seeding exercise_records"

user = User.first
ExerciseRecord.find_or_create_by!(
  user_id: user.id,
  recorded_date: Date.today,
  content: "ランニング30分",
  exercise_time: 30
)

ExerciseRecord.find_or_create_by!(
  user_id: user.id,
  recorded_date: Date.today - 1,
  content: "筋トレ（腕立て・腹筋）20分",
  exercise_time: 20
)


puts "Seeding characters..."

Character.find_or_create_by!(
  name: "ランナー",
  base_image: "runner_base.png"
)

Character.find_or_create_by!(
  name: "マッスル",
  base_image: "muscle_base.png"
)





puts "Seeding user_characters..."

user = User.first
character = Character.first

if user && character
  UserCharacter.find_or_create_by!(
    user_id: user.id,
    character_id: character.id
  ) do |uc|
    uc.level = 1
    uc.current_points = 50
    uc.require_points = 100
    uc.character_animation = "walk"
  end

  puts "✅ UserCharacter created for user_id=#{user.id}, character_id=#{character.id}"
else
  puts "⚠️ User または Character のデータが存在しないため、user_characters は作成されませんでした"
end