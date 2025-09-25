Rails.application.routes.draw do
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Reveal health status on /up that returns 200 if the app boots with no exceptions, otherwise 500.
  # Can be used by load balancers and uptime monitors to verify that the app is live.
  get "up" => "rails/health#show", as: :rails_health_check

  # Defines the root path route ("/")
  # root "posts#index"

  resources :exercise_records, only: [:create, :index, :show, :update, :destroy]
  namespace :api do
    namespace :v1 do
      resources :exercise_records do
        # 運動記録にいいね
        resource :like, only: [:create, :destroy]
      end
      # 新規登録
      post "users", to: "users#create"
      # ログイン
      post "login", to: "sessions#create"
    end
  end
end
