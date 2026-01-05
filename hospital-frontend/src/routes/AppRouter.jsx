import React from "react";
import { Routes, Route, Navigate } from "react-router-dom";

/* Layouts */
import MainLayout from "../layout/MainLayout";
import DashboardLayout from "../layout/DashboardLayout";

/* Páginas públicas */
import Home from "../pages/Home";
import Servicios from "../pages/Servicios";
import Clinicas from "../pages/Clinicas";
import Transparencia from "../pages/Transparencia";
import Contacto from "../pages/Contacto";

/* Dashboard */
import DashboardHome from "../pages/dashboard/DashboardHome";
import Doctores from "../pages/dashboard/doctores/Doctores";
import Pacientes from "../pages/dashboard/pacientes/Pacientes";
import Citas from "../pages/dashboard/citas/Citas";
import Consultas from "../pages/dashboard/consultas/Consultas";

/* Rutas protegidas */
import ProtectedRoute from "./ProtectedRoute";
import Login from "../pages/auth/Login";

export default function AppRouter() {
  return (
    <Routes>

      {/* ===================== */}
      {/* RUTAS PÚBLICAS */}
      {/* ===================== */}
      <Route element={<MainLayout />}>
        <Route path="/" element={<Home />} />
        <Route path="/servicios" element={<Servicios />} />
        <Route path="/clinicas" element={<Clinicas />} />
        <Route path="/transparencia" element={<Transparencia />} />
        <Route path="/contacto" element={<Contacto />} />
        <Route path="/login" element={<Login />} />
      </Route>

      {/* ===================== */}
      {/* DASHBOARD */}
      {/* ===================== */}
      <Route path="/dashboard" element={<DashboardLayout />}>
        <Route index element={<DashboardHome />} />
        <Route path="doctores" element={<Doctores />} />
        <Route path="pacientes" element={<Pacientes />} />
        <Route path="citas" element={<Citas />} />
        <Route path="consultas" element={<Consultas />} />
      </Route>

      {/* ===================== */}
      {/* RUTAS PROTEGIDAS */}
      {/* ===================== */}
      <Route element={<ProtectedRoute />}>
        <Route path="/dashboard" element={<DashboardLayout />} />
      </Route>

      {/* ===================== */}
      {/* FALLBACK */}
      {/* ===================== */}
      <Route path="*" element={<Navigate to="/" replace />} />

    </Routes>
  );
}
