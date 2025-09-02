const express = require("express");
const app = express();
const PORT = 3000;

app.use(express.json());

const clientes = [
  {
    id: 1,
    nome: "Maria Silva",
    email: "maria.silva@email.com",
    telefone: "11987654321",
    empresa: "Tech Solutions",
  },
  {
    id: 2,
    nome: "João Oliveira",
    email: "joao.oliveira@email.com",
    telefone: "21912345678",
    empresa: "Inova Consultoria",
  },
  {
    id: 3,
    nome: "Ana Souza",
    email: "ana.souza@email.com",
    telefone: "31998765432",
    empresa: "Design Criativo",
  },
];

app.get("/clientes", (req, res) => {
  res.json(clientes);
});

app.get("/clientes/:id", (req, res) => {
  const cliente = clientes.find((c) => c.id === parseInt(req.params.id));
  if (cliente) {
    res.json(cliente);
  } else {
    res.status(404).json({ message: "Cliente não encontrado" });
  }
});

app.post("/clientes", (req, res) => {
  const novoCliente = {
    id: clientes.length + 1,
    ...req.body,
  };
  clientes.push(novoCliente);
  res.status(201).json(novoCliente);
});

app.put("/clientes/:id", (req, res) => {
  const clienteIndex = clientes.findIndex(
    (c) => c.id === parseInt(req.params.id)
  );
  if (clienteIndex !== -1) {
    clientes[clienteIndex] = { id: parseInt(req.params.id), ...req.body };
    res.json(clientes[clienteIndex]);
  } else {
    res.status(404).json({ message: "Cliente não encontrado" });
  }
});

app.delete("/clientes/:id", (req, res) => {
  const clienteIndex = clientes.findIndex(
    (c) => c.id === parseInt(req.params.id)
  );
  if (clienteIndex !== -1) {
    const clienteRemovido = clientes.splice(clienteIndex, 1);
    res.json(clienteRemovido[0]);
  } else {
    res.status(404).json({ message: "Cliente não encontrado" });
  }
});

app.listen(PORT, () => {
  console.log(`Servidor rodando em http://localhost:${PORT}`);
});
