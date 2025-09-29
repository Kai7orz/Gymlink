if Rails.env.production?
  $redis = Redis.new(url: ENV.fetch("REDIS_URL") { "redis://localhost:6379/1" })
end