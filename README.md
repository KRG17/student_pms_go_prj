# 💡 Student Project Management System

The Student Project Management System empowers students to propose, update, and withdraw project submissions, contributing to a dynamic project marketplace. Mentors gain visibility into all student-entered projects, fostering collaboration and guidance. Implemented in Go with GoFr framework and SQL, the system showcases CRUD operations, real-time collaboration, and seamless database integration.

## API Requests made using postman

## 1.🚀 GET Request
- Get all the projects
  ![Screenshot (54)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/59c15d0c-da5a-4364-80e0-2811e857b5ea)

## 2.➕ POST Request
- Add a new project
  ![Screenshot (55)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/4e2c85ef-75a2-4af9-982c-0b13a3a3a520)

- After adding a new project
  ![Screenshot (56)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/f7fe27ac-0c41-459a-8445-626ed6c3ddf9)


## 3.🔄 PUT Request
- Updates an existing project
![Screenshot (57)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/49dfb097-0c9b-41ca-877f-10267272b9d0)

## 4.🗑️ DELETE Request
- Deletes an existing project
![Screenshot (58)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/7faecf19-4dc8-41b5-ba36-f4242d8d8caf)

# UML Diagrams
- Use Case Diagram
  
  ![usecase-diagram-student-pms](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/33158205-590e-4ac3-bfaa-00a36553dab4)
  
- Class Diagram
  ![class-diagram-student-pms](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/9501eb1a-962b-4d2e-9401-cfe371f2d22a)

- Sequence Diagram
 ![sequence-diagram-student-pms](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/c63ac020-c8ae-4429-92fa-0970bb4781ed)

# Database 🛢️

The project utilizes a MySQL database to store information about student projects. The MySQL database is set up as a Docker container for easy deployment and management.

## Database Schema 🗃️
- Table: `projects`
  - Columns: `id` (int), `name` (varchar), `description` (text)
![Screenshot (59)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/a231f36d-7303-4730-ba43-e3fd51f2d576)


## Database Setup 🐳
- Docker Container
![Screenshot (60)](https://github.com/KRG17/zopsmart_go_prj/assets/109519365/58c84e1d-d7e1-46c3-84e0-64cf03dbac8a)

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
