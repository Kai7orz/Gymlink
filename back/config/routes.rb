Rails.application.routes.draw do
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Reveal health status on /up that returns 200 if the app boots with no exceptions, otherwise 500.
  # Can be used by load balancers and uptime monitors to verify that the app is live.
  get "up" => "rails/health#show", as: :rails_health_check

  # Defines the root path route ("/")
  # root "posts#index"

  # 新規登録
  post "users", to: "users#create"

  # ログイン
  post "login", to: "sessions#create"

  # プロフィール表示
  get "user_profiles/:user_id", to: "users#profile"

  # 運動記録機能
  get "users/:user_id/exercises", to: "exercise_records#user_exercises"
  resources :exercises, only: [:create, :index], controller: 'exercise_records'

  # いいね機能
  post "likes", to: "likes#create"
  delete "likes/:exercise_record_id", to: "likes#destroy"

  # チーム機能
  get "teams", to: "teams#show"

  # フォロー機能
  post "follows", to: "follows#create"
  post "unfollows", to: "follows#destroy"

  # キャラクター機能
  get "user_characters/:user_id", to: "user_characters#show"
end
