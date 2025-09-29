FirebaseIdToken.configure do |config|
  config.redis = Redis.new(url: ENV.fetch("REDIS_URL", "redis://127.0.0.1:6379/1"))
  config.project_ids = ["rizap-hackathon"] # FirebaseプロジェクIDD
end

# 起動時に証明書を取得してRedisへキャッシュ
begin
  FirebaseIdToken::Certificates.request
rescue => e
  Rails.logger.warn("[firebase_id_token] warmup failed: #{e.class}: #{e.message}")
end