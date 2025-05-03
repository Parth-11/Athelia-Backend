# Athelia-Backend

# ModLib

ModLib is a modular Go-based application that demonstrates a clean separation of concerns across client, server, and internal libraries. It includes CLI tools, authentication logic, and reusable packages.

---

## 📁 Project Structure

├── LICENSE # License file
├── README.md # Project documentation
├── config.go # Application-wide configuration
├── go.mod # Go module definition
├── go.sum # Go module checksums
├── clientlib/ # Reusable client library
│ ├── lib.go # Core client functions
│ └── lib_test.go # Tests for clientlib
├── cmd/ # Main entry points
│ ├── modlib-client/ # CLI for client-side interaction
│ │ └── main.go # Entry point for client CLI
│ └── modlib-server/ # CLI for server-side app
│ └── main.go # Entry point for server CLI
├── internal/ # Private internal packages
│ └── auth/ # Authentication utilities
│ ├── auth.go # Auth implementation
| └── auth_test.go # Tests for auth
└── serverlib/ # Server-side reusable logic
  └── lib.go # Core server functions