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





users = [
  User.find_or_create_by!(firebase_uid: "uid_001") do |u|
    u.name = "Taro Yamada"
    u.color = "#ff0000"
    u.avatar_url = "https://example.com/avatar1.png"
    u.team_id = dummy_team.id
  end,
  User.find_or_create_by!(firebase_uid: "uid_002") do |u|
    u.name = "Hanako Suzuki"
    u.color = "#00ff00"
    u.avatar_url = "https://example.com/avatar2.png"
    u.team_id = dummy_team.id
  end,
  User.find_or_create_by!(firebase_uid: "uid_003") do |u|
    u.name = "Tomoko Yamamoto"
    u.color = "red"
    u.avatar_url = "https://example.com/avatar3.png"
    u.team_id = dummy_team.id
  end
]



puts "Seeding exercise_records"

exercise_records = 10.times.map do |i|
  ExerciseRecord.create!(
    user_id: users.sample.id,
    exercise_time: rand(10..120),
    content: "運動内容#{i + 1}",
    recorded_date: Date.today - i.days
  )
end

puts "✅ ExerciseRecords created: #{exercise_records.size}"



puts "Seeding characters..."

Character.find_or_create_by!(
  name: "ランナー",
  base_image: "runner_base.png"
)

Character.find_or_create_by!(
  name: "マッスル",
  base_image: "muscle_base.png"
)




puts "Seeding user_characters"


user = users.first
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


Profile.find_or_create_by!(
  user_id: users[0].id,
  name: "山田太郎",
  profile_image: "/images/profile1.png"
)
Profile.find_or_create_by!(
  user_id: users[1].id,
  name: "佐藤花子",
  profile_image: "/images/profile2.png"
)
Profile.find_or_create_by!(
  user_id: users[2].id,
  name: "山本ともこ",
  profile_image: "/images/profile3.png"
)




# User1 → User2, User3 をフォロー
Follow.find_or_create_by!(follower_id: users[0].id, followed_id: users[1].id)
Follow.find_or_create_by!(follower_id: users[0].id, followed_id: users[2].id)

# User2 → User1 をフォロー
Follow.find_or_create_by!(follower_id: users[1].id, followed_id: users[0].id)

# User3 → User1, User2 をフォロー
Follow.find_or_create_by!(follower_id: users[2].id, followed_id: users[0].id)
Follow.find_or_create_by!(follower_id: users[2].id, followed_id: users[1].id)


puts "✅ Follows created!"

20.times do
  liker = users.sample
  record = exercise_records.sample
  UserLike.find_or_create_by!(user_id: liker.id, exercise_record_id: record.id)
end

puts "✅ UserLike dummy data created!"




puts "Seeding character_growths..."
Character.all.each do |character|
  (1..5).each do |level|
    CharacterGrowth.find_or_create_by!(
      character_id: character.id,
      level: level
    ) do |growth|
      growth.required_points = level * 100
      growth.reward_image_url = "/images/#{character.name.downcase}_level#{level}.png"
    end
  end
end

puts "✅ CharacterGrowths created!"