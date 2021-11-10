# Dataproduct Topics
## Prerequisites
```
gcloud iam --project aura-prod-d7e3 \
  service-accounts create dataproduct-topics \
  --description="Manually created service-account for dataproduct-topics"

bq mk \
  --project_id aura-prod-d7e3 \
  --location europe-north1 \
  dataproduct_topics.dataproduct_topics \
  cluster:string,topic:string,team:string

bq add-iam-policy-binding \
  --member='serviceAccount:dataproduct-topics@aura-prod-d7e3.iam.gserviceaccount.com' \
  --role='roles/bigquery.dataEditor' \
  aura-prod-d7e3:dataproduct_topics.dataproduct_topics
```
