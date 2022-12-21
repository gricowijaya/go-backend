-- name: CreateUser :one
INSERT INTO "Users" (
    email,
    username,
    password
    ) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetAllUsers :many 
SELECT * FROM "Users";

-- name: GetOneUser :one 
SELECT * FROM "Users" WHERE "Users"."id" = $1 LIMIT 1;

-- name: UpdateEmail :exec 
UPDATE "Users" SET "email" = $2 WHERE "id" = $1;

-- name: DeleteUser :exec
DELETE FROM "Users" WHERE "id" = $1;
