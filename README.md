# golang_ecommerce_api

A simple e-commerce web application for managing products, categories, comments, orders, and user accounts.

## Description

This project is a web-based e-commerce application designed to handle various aspects of online shopping. It includes features such as product management,
category management, user registration and authentication, comment posting, and order management. It's a versatile platform for both customers and administrators to interact with the online store.

Usage
To use this application, you can interact with the following routes:

Categories
POST /category: Add a new category.
GET /category: Get all categories.
Comments
POST /comment: Add a new comment.
GET /comment: Get all comments.
GET /comment/:id: Get a comment by ID.
DELETE /comment/:id: Delete a comment by ID.
GET /comment/product-id/:id: Get comments by product ID.
GET /comment/user-id/:id: Get comments by user ID.
Orders
POST /order: Add a new order.
Products
GET /products: Get all products (authentication required).
GET /products/:id: Get a product by ID.
POST /products: Add a new product (admin authentication required).
DELETE /products/:id: Delete a product by ID.
PUT /products/:id: Update a product by ID.
Users
POST /user/register: Register a new user.
POST /user/login: User login.
Models
Here are the models used in this project:

Category
Category: Represents a category with properties Category and ID.
Comments
Comments: Represents a comment with properties ProductID, CustomerID, and CommentText.
Orders
Orders: Represents an order with properties CustomerID, OrderDate, Status, ProductID, and Quantity.
Product
Product: Represents a product with properties Name, Description, Price, and CategoryID.
User
User: Represents a user with properties UserName, Email, Password, Phone, Role, Disabled, and DateOfBirth.
