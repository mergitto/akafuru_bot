gcloud functions deploy "$FUNCTION_NAME" \
--entry-point AkafuruCommand \
--runtime go113 \
--region asia-northeast1 \
--trigger-http \
--set-env-vars VERIFICATION_TOKEN="$VERIFICATION_TOKEN",API_KEY="$API_KEY",BLOG_ACTIVITY_API="$BLOG_ACTIVITY_API"
