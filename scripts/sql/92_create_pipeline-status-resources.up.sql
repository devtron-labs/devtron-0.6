CREATE SEQUENCE IF NOT EXISTS id_seq_pipeline_status_timeline_resources;

CREATE TABLE public.pipeline_status_timeline_resources (
"id"                                integer NOT NULL DEFAULT nextval('id_seq_pipeline_status_timeline_resources'::regclass),
"installed_app_version_history_id"  integer,
"cd_workflow_runner_id"             integer,
"resource_name"                     text,
"resource_kind"                     text,
"resource_group"                    text,
"resource_phase"                    text,
"resource_status"                   text,
"status_message"                    text,
"timeline_stage"                    text DEFAULT 'KUBECTL_APPLY',
"created_on"                        timestamptz,
"created_by"                        int4,
"updated_on"                        timestamptz,
"updated_by"                        int4,
CONSTRAINT "pipeline_status_timeline_resources_cd_workflow_runner_id_fkey" FOREIGN KEY ("cd_workflow_runner_id") REFERENCES "public"."cd_workflow_runner" ("id"),
CONSTRAINT "pipeline_status_timeline_resources_installed_app_version_history_id_fkey" FOREIGN KEY ("installed_app_version_history_id") REFERENCES "public"."installed_app_version_history" ("id"),
PRIMARY KEY ("id")
);


CREATE SEQUENCE IF NOT EXISTS id_seq_pipeline_status_fetch_detail;

CREATE TABLE public.pipeline_status_fetch_detail (
"id"                                integer NOT NULL DEFAULT nextval('id_seq_pipeline_status_fetch_detail'::regclass),
"installed_app_version_history_id"  integer,
"cd_workflow_runner_id"             integer,
"last_fetched_at"                   timestampz,
"fetch_count"                       integer,
"created_on"                        timestamptz,
"created_by"                        int4,
"updated_on"                        timestamptz,
"updated_by"                        int4,
 CONSTRAINT "pipeline_status_fetch_detail_cd_workflow_runner_id_fkey" FOREIGN KEY ("cd_workflow_runner_id") REFERENCES "public"."cd_workflow_runner" ("id"),
 CONSTRAINT "pipeline_status_fetch_detail_installed_app_version_history_id_fkey" FOREIGN KEY ("installed_app_version_history_id") REFERENCES "public"."installed_app_version_history" ("id"),
 PRIMARY KEY ("id")
);