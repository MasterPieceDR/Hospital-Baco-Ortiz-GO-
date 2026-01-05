// src/dashboard/pages/Doctores.jsx
import React, { useEffect, useState } from "react";
import api from "../../api/axios";

export default function Doctores() {
	const [doctores, setDoctores] = useState([]);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState(null);

	const [showModal, setShowModal] = useState(false);
	const [editing, setEditing] = useState(null); // medico object when editing
	const [form, setForm] = useState({ nombre: "", apellido: "", correo: "", telefono: "", activo: true });

	const fetch = () => {
		setLoading(true);
		setError(null);
		api.get("/medicos")
			.then(r => setDoctores(r.data || []))
			.catch(e => setError(e.response?.data?.error || e.message))
			.finally(() => setLoading(false));
	};

	useEffect(() => { fetch(); }, []);

	const openCreate = () => { setEditing(null); setForm({ nombre: "", apellido: "", correo: "", telefono: "", activo: true }); setShowModal(true); };
	const openEdit = (m) => { setEditing(m); setForm({ nombre: m.nombre || m.Nombre || "", apellido: m.apellido || m.Apellido || "", correo: m.correo || m.email || "", telefono: m.telefono || m.Telefono || "", activo: m.activo===false?false:true }); setShowModal(true); };

	const save = async () => {
		try {
			if (editing) {
				await api.put(`/medicos/${editing.id || editing.medico_id}`, {
					nombre: form.nombre, apellido: form.apellido, correo: form.correo, telefono: form.telefono, activo: form.activo
				});
			} else {
				await api.post("/medicos", {
					nombre: form.nombre, apellido: form.apellido, correo: form.correo, telefono: form.telefono, activo: true
				});
			}
			setShowModal(false);
			fetch();
		} catch (e) {
			alert(e.response?.data?.error || e.message);
		}
	};

	const remove = async (m) => {
		if (!confirm("Eliminar médico?")) return;
		try {
			await api.delete(`/medicos/${m.id || m.medico_id}`);
			fetch();
		} catch (e) {
			alert(e.response?.data?.error || e.message);
		}
	};

	if (loading) return <div style={{ padding: 24 }}>Cargando médicos...</div>;
	if (error) return <div style={{ padding: 24, color: "red" }}>{error}</div>;

	return (
		<div style={{ padding: 24 }}>
			<div style={{ display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: 16 }}>
				<h2>Gestión de Doctores</h2>
				<div>
					<button className="btn" onClick={openCreate}>Nuevo Médico</button>
					<button className="btn secondary" style={{ marginLeft: 8 }} onClick={fetch}>Refrescar</button>
				</div>
			</div>

			{doctores.length === 0 ? <div className="card">(sin médicos)</div> : (
				<table className="table">
					<thead>
						<tr><th>Nombre</th><th>Email</th><th>Teléfono</th><th>Activo</th><th></th></tr>
					</thead>
					<tbody>
						{doctores.map(d => (
							<tr key={d.id || d.medico_id}>
								<td>{(d.nombre || d.Nombre)} {(d.apellido||d.Apellido||"")}</td>
								<td>{d.correo || d.email || d.Correo || "-"}</td>
								<td>{d.telefono || d.Telefono || "-"}</td>
								<td>{d.activo === false ? "No" : "Sí"}</td>
								<td style={{ width: 220 }}>
									<div style={{ display: "flex", justifyContent: "flex-end", gap: 8 }}>
										<button className="btn secondary" onClick={() => openEdit(d)}>Editar</button>
										<button className="btn danger" onClick={() => remove(d)}>Eliminar</button>
									</div>
								</td>
							</tr>
						))}
					</tbody>
				</table>
			)}

			{showModal && (
				<div className="modal-backdrop" onClick={() => setShowModal(false)}>
					<div className="modal" onClick={(e)=>e.stopPropagation()}>
						<h3>{editing ? "Editar Médico" : "Crear Médico"}</h3>
						<div className="form-row">
							<label>Nombre</label>
							<input value={form.nombre} onChange={e=>setForm({...form,nombre:e.target.value})} />
						</div>
						<div className="form-row">
							<label>Apellido</label>
							<input value={form.apellido} onChange={e=>setForm({...form,apellido:e.target.value})} />
						</div>
						<div className="form-row">
							<label>Correo</label>
							<input value={form.correo} onChange={e=>setForm({...form,correo:e.target.value})} />
						</div>
						<div className="form-row">
							<label>Teléfono</label>
							<input value={form.telefono} onChange={e=>setForm({...form,telefono:e.target.value})} />
						</div>
						<div className="actions">
							<button className="btn" onClick={save}>{editing ? "Guardar" : "Crear"}</button>
							<button className="btn secondary" onClick={()=>setShowModal(false)}>Cancelar</button>
						</div>
					</div>
				</div>
			)}
		</div>
	);
}
