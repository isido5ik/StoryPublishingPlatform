
CREATE TABLE t_users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE t_admin(
    admin_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES t_users(user_id) ON DELETE CASCADE
);

CREATE TABLE t_client(
    client_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES t_users(user_id) ON DELETE CASCADE
);

CREATE TABLE t_roles(
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(100) NOT NULL
);

CREATE TABLE t_users_roles (
    user_id INTEGER REFERENCES t_users(user_id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES t_roles(role_id) ON DELETE CASCADE
);


CREATE TABLE t_posts (
    post_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES t_users(user_id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES t_categories(category_id) ON DELETE SET NULL,
    title TEXT NOT NULL, 
    content TEXT NOT NULL,
    comments INTEGER,
    likes INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE t_categories (
    category_id SERIAL PRIMARY KEY,
    category_name VARCHAR(100) NOT NULL
);

CREATE TABLE t_comments (
    comment_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES t_users(user_id) ON DELETE CASCADE,
    post_id INTEGER REFERENCES t_posts(post_id) ON DELETE CASCADE,
    parent_id INTEGER REFERENCES t_comments(comment_id) ON DELETE CASCADE,
    comment_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE t_likes (
    like_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES t_users(user_id) ON DELETE CASCADE,
    post_id INTEGER REFERENCES t_posts(post_id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
