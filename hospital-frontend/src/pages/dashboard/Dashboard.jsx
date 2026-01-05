import React, { useEffect, useState, useContext } from "react";
import api from "../../api/axios";
import { AuthContext } from "../../context/AuthContext";

export default function Dashboard() {
  const [medicos, setMedicos] = useState([]);
  const [err, setErr] = useState(null);
  const { logout, user } = useContext(AuthContext);

  useEffect(() => {
    api.get("/medicos")
      .then(r => setMedicos(r.data || []))
      .catch(e => setErr(e.response?.data?.error || e.message));
  }, []);

  return (
    <div style={{ padding: 20 }}>
      <h2>Dashboard</h2>
      <div>
        Usuario: {user?.username || "Anónimo"} — <button onClick={() => { logout(); window.location.href = "/"; }}>Cerrar sesión</button>
      </div>
      {err && <div style={{ color: "red" }}>{err}</div>}
      <h3>Médicos</h3>
      <ul>
        {medicos.length === 0 && <li>(sin médicos o sin datos)</li>}
        {medicos.map(m => <li key={m.id || m.medico_id}>{m.nombre || m.Nombre || `${m.Nombre} ${m.Apellido}`}</li>)}
      </ul>
    </div>
  );
}
