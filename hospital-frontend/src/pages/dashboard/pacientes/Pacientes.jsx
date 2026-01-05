import React, { useEffect, useState } from "react";
import api from "../../../api/axios";

export default function Pacientes() {
	const [list, setList] = useState([]);
	const [loading, setLoading] = useState(true);
	const [showForm, setShowForm] = useState(false);
	const [editing, setEditing] = useState(null);
	const [form, setForm] = useState({ nombre: "", apellido: "", correo: "", telefono: "" });

	const fetch = () => {
		setLoading(true);
		api.get("/pacientes")
			.then(r => setList(r.data || []))
			.catch(e => console.error(e))
			.finally(()=>setLoading(false));
	};

	useEffect(()=>fetch(),[]);

	const openCreate = () => { setEditing(null); setForm({ nombre:"", apellido:"", correo:"", telefono:"" }); setShowForm(true); };
	const openEdit = (p) => { setEditing(p); setForm({ nombre:p.nombre||p.Nombre, apellido:p.apellido||p.Apellido, correo:p.correo||p.email||"", telefono:p.telefono||p.Telefono||"" }); setShowForm(true); };

	const save = async () => {
		try {
			if (editing) await api.put(`/pacientes/${editing.id || editing.paciente_id}`, form);
			else await api.post("/pacientes", form);
			setShowForm(false); fetch();
		} catch (e) { alert(e.response?.data?.error||e.message); }
	};

	const remove = async (p) => {
		if(!confirm("Eliminar paciente?")) return;
		await api.delete(`/pacientes/${p.id || p.paciente_id}`);
		fetch();
	};

	if (loading) return <div style={{padding:24}}>Cargando pacientes...</div>;

	return (
		<div style={{padding:24}}>
			<div style={{display:"flex",justifyContent:"space-between",alignItems:"center"}}>
				<h2>Gestión de Pacientes</h2>
				<div>
					<button className="btn" onClick={openCreate}>Nuevo Paciente</button>
				</div>
			</div>

			{list.map(p => {
				// intentar diferentes nombres de campo que pueden venir del backend
				const ci = p.cedula || p.numero_seguridad_social || p.numeroSeguridadSocial || p.nss || "";
				return (
					<div className="card" key={p.id || p.paciente_id}>
						<div className="list-row">
							<div>
								<strong>{p.nombre || p.Nombre}</strong>
								<div style={{color:"#666"}}>CI: {ci}</div>
							</div>
							<div style={{display:"flex",gap:8}}>
								<button className="btn secondary" onClick={()=>openEdit(p)}>Editar</button>
								<button className="btn danger" onClick={()=>remove(p)}>Eliminar</button>
							</div>
						</div>
					</div>
				);
			})}

			{showForm && (
				<div className="modal-backdrop" onClick={()=>setShowForm(false)}>
					<div className="modal" onClick={e=>e.stopPropagation()}>
						<h3>{editing ? "Editar paciente" : "Crear paciente"}</h3>
						<div className="form-row"><label>Nombre</label><input value={form.nombre} onChange={e=>setForm({...form,nombre:e.target.value})} /></div>
						<div className="form-row"><label>Apellido</label><input value={form.apellido} onChange={e=>setForm({...form,apellido:e.target.value})} /></div>
						<div className="form-row"><label>Email</label><input value={form.correo} onChange={e=>setForm({...form,correo:e.target.value})} /></div>
						<div className="form-row"><label>Teléfono</label><input value={form.telefono} onChange={e=>setForm({...form,telefono:e.target.value})} /></div>
						<div className="actions"><button className="btn" onClick={save}>Guardar</button><button className="btn secondary" onClick={()=>setShowForm(false)}>Cancelar</button></div>
					</div>
				</div>
			)}
		</div>
	);
}
