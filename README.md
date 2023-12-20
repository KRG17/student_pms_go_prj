# 💡 Student Project Management System

The Student Project Management System empowers students to propose, update, and withdraw project submissions, contributing to a dynamic project marketplace. Mentors gain visibility into all student-entered projects, fostering collaboration and guidance. Implemented in Go with GoFr framework and SQL, the system showcases CRUD operations, real-time collaboration, and seamless database integration.

## API Requests made using postman

## 1.🚀 GET Request
- Get all the projects
  
![Screenshot (54)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/2c857c84-78a7-4187-8a43-f4e8a4aaca97)

## 2.➕ POST Request
- Add a new project
  
![Screenshot (55)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/bf3d7ff1-1b26-42ad-8261-c218b8be34f5)

- After adding a new project
  
![Screenshot (61)](https://github.com/KRG17/zopsmart_GoFr_api_project/assets/109519365/9d510ac9-387d-4374-8191-80e351e2e0b6)

## 3.🔄 PUT Request
- Updates an existing project
  
![Screenshot (57)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/0856901a-ac23-461c-96b7-5d76bc63ac4b)

## 4.🗑️ DELETE Request
- Deletes an existing project
  
![Screenshot (58)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/e1855310-f7e9-4eb1-ac26-50f88c901faa)

# UML Diagrams
- Use Case Diagram
  
  ![usecase-diagram-student-pms](https://github.com/KRG17/zopsmart_GoFr_api_project/assets/109519365/280269b1-95b6-4d2b-92e7-8ce82d5b35f2)

- Class Diagram
  
  ![class-diagram-student-pms2](https://github.com/KRG17/zopsmart_GoFr_api_project/assets/109519365/c5ad5c2b-a6a9-4bae-9bd6-b246b4f30db6)

- Sequence Diagram
  
  ![sequence-diagram-student-pms2](https://github.com/KRG17/zopsmart_GoFr_api_project/assets/109519365/1bf83ef4-95e8-4401-8a64-ed9e4f8ee90e)

# Database 🛢️

The project utilizes a MySQL database to store information about student projects. The MySQL database is set up as a Docker container for easy deployment and management.

## Database Schema 🗃️

- Table: `projects`
  - Columns: `id` (int), `name` (varchar), `description` (text)
![Screenshot (59)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/d68758bb-1dc6-49c1-8bff-c2f1b02e0dda)


## Database Setup 🐳
- Docker Container
![Screenshot (60)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/07fd2a83-d48d-4755-a8ce-a51a3835bcf8)

- To run the MySQL database locally as a Docker container, use the following command:
```bash
docker run --name student-project-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=student_project_db -p 3306:3306 -d mysql:8.0.30
```
- After this create a new connection in MYSQL workbench and make a projects table with the above defined schema to successfully setup the database.

# ⚡️ RUN
- Now simply run the following command in your terminal to run the application:
```bash
go run cmd/main.go
```
