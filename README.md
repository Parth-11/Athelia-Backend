# Athelia-Backend

# ModLib

ModLib is a modular Go-based application that demonstrates a clean separation of concerns across client, server, and internal libraries. It includes CLI tools, authentication logic, and reusable packages.

---

## 📁 Project Structure
```
.
├── LICENSE
├── README.md
├── config.go
├── go.mod
├── go.sum
├── clientlib/
│   ├── lib.go
│   └── lib_test.go
├── cmd/
│   ├── modlib-client/
│   │   └── main.go
│   └── modlib-server/
│       └── main.go
├── internal/
│   └── auth/
│       ├── auth.go
│       └── auth_test.go
└── serverlib/
    └── lib.go
```