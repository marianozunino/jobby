-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "email" character varying NOT NULL, "password" character varying NOT NULL, "role" character varying NOT NULL, "reset_password_token" character varying NULL, "reset_password_expires" timestamptz NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "full_name" character varying NOT NULL, "salt" character varying NOT NULL, "external_id" integer NULL, PRIMARY KEY ("id"));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create "contact_us_messages" table
CREATE TABLE "contact_us_messages" ("id" uuid NOT NULL, "name" character varying NOT NULL, "email" character varying NOT NULL, "message" character varying NOT NULL, "phone" character varying NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create "applicant_profiles" table
CREATE TABLE "applicant_profiles" ("id" uuid NOT NULL, "birthday" timestamptz NOT NULL, "gender" character varying NOT NULL, "phone" character varying NOT NULL, "address1" character varying NOT NULL, "address2" character varying NOT NULL, "cv" character varying NULL, "internal_comments" character varying NOT NULL, "receive_emails" boolean NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "extra_skills" jsonb NULL, "user_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "applicant_profiles_users_applicant_profiles" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "categories" table
CREATE TABLE "categories" ("id" uuid NOT NULL, "name" character varying NOT NULL, "slug" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "is_root" boolean NOT NULL, "parent_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "categories_categories_child_categories" FOREIGN KEY ("parent_id") REFERENCES "categories" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "status" table
CREATE TABLE "status" ("id" uuid NOT NULL, "name" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create "job_offers" table
CREATE TABLE "job_offers" ("id" uuid NOT NULL, "title" character varying NOT NULL, "reference" integer NOT NULL, "start_date" timestamptz NOT NULL, "end_date" timestamptz NOT NULL, "address1" character varying NULL, "address2" character varying NULL, "department" character varying NOT NULL, "description" character varying NOT NULL, "working_hours" character varying NOT NULL, "salary" character varying NOT NULL, "slug" character varying NOT NULL, "is_featured" boolean NULL, "has_been_emailed" boolean NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "status_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "job_offers_status_job_offers" FOREIGN KEY ("status_id") REFERENCES "status" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "job_offers_slug_key" to table: "job_offers"
CREATE UNIQUE INDEX "job_offers_slug_key" ON "job_offers" ("slug");
-- Create "job_offer_categories" table
CREATE TABLE "job_offer_categories" ("id" uuid NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "category_id" uuid NOT NULL, "job_offer_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "job_offer_categories_categories_job_offer_categories" FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "job_offer_categories_job_offers_job_offer_categories" FOREIGN KEY ("job_offer_id") REFERENCES "job_offers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "posts" table
CREATE TABLE "posts" ("id" uuid NOT NULL, "title" character varying NOT NULL, "content" character varying NOT NULL, "slug" character varying NOT NULL, "is_highlighted" boolean NOT NULL, "is_published" boolean NOT NULL, "published_at" timestamptz NULL, "preview_image" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "author_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "posts_users_posts" FOREIGN KEY ("author_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "posts_slug_key" to table: "posts"
CREATE UNIQUE INDEX "posts_slug_key" ON "posts" ("slug");
-- Create "applicant_interests" table
CREATE TABLE "applicant_interests" ("id" uuid NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "applicant_profile_id" uuid NOT NULL, "category_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "applicant_interests_applicant_profiles_applicant_interests" FOREIGN KEY ("applicant_profile_id") REFERENCES "applicant_profiles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "applicant_interests_categories_applicant_interests" FOREIGN KEY ("category_id") REFERENCES "categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "applications" table
CREATE TABLE "applications" ("id" uuid NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "applicant_profile_id" uuid NOT NULL, "job_offer_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "applications_applicant_profiles_applications" FOREIGN KEY ("applicant_profile_id") REFERENCES "applicant_profiles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "applications_job_offers_applications" FOREIGN KEY ("job_offer_id") REFERENCES "job_offers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "degree_levels" table
CREATE TABLE "degree_levels" ("id" uuid NOT NULL, "name" character varying NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create "education" table
CREATE TABLE "education" ("id" uuid NOT NULL, "title" character varying NOT NULL, "institution" character varying NULL, "date_obtained" timestamptz NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "comments" character varying NULL, "applicant_profile_id" uuid NULL, "applicant_profile_educations" uuid NULL, "degree_level_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "education_applicant_profiles_educations" FOREIGN KEY ("applicant_profile_educations") REFERENCES "applicant_profiles" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "education_degree_levels_educations" FOREIGN KEY ("degree_level_id") REFERENCES "degree_levels" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "interviews" table
CREATE TABLE "interviews" ("id" uuid NOT NULL, "comment" character varying NOT NULL, "interview_date" timestamptz NOT NULL, "status" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "application_id" uuid NULL, "interviewer_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "interviews_applications_interviews" FOREIGN KEY ("application_id") REFERENCES "applications" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "interviews_users_interviews" FOREIGN KEY ("interviewer_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "languages" table
CREATE TABLE "languages" ("id" serial NOT NULL, "name" character varying NOT NULL, "level" character varying NOT NULL, "comments" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "applicant_profile_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "languages_applicant_profiles_languages" FOREIGN KEY ("applicant_profile_id") REFERENCES "applicant_profiles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "post_categories" table
CREATE TABLE "post_categories" ("id" uuid NOT NULL, "name" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create index "post_categories_name_key" to table: "post_categories"
CREATE UNIQUE INDEX "post_categories_name_key" ON "post_categories" ("name");
-- Create "post_category" table
CREATE TABLE "post_category" ("category_id" uuid NOT NULL, "post_id" uuid NOT NULL, PRIMARY KEY ("category_id", "post_id"), CONSTRAINT "post_category_category_id" FOREIGN KEY ("category_id") REFERENCES "post_categories" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "post_category_post_id" FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "work_experience" table
CREATE TABLE "work_experience" ("id" uuid NOT NULL, "company" character varying NULL, "position" character varying NULL, "description" character varying NULL, "start_date" timestamptz NULL, "end_date" timestamptz NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "internal_comments" character varying NULL, "applicant_profile_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "work_experience_applicant_profiles_work_experiences" FOREIGN KEY ("applicant_profile_id") REFERENCES "applicant_profiles" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "skills" table
CREATE TABLE "skills" ("id" uuid NOT NULL, "name" character varying NOT NULL, "description" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create "applicant_profile_skills" table
CREATE TABLE "applicant_profile_skills" ("id" uuid NOT NULL, "level" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "applicant_profile_id" uuid NOT NULL, "skill_id" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "applicant_profile_skills_applicant_profiles_applicant_profile_s" FOREIGN KEY ("applicant_profile_id") REFERENCES "applicant_profiles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "applicant_profile_skills_skills_applicant_profile_skills" FOREIGN KEY ("skill_id") REFERENCES "skills" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);