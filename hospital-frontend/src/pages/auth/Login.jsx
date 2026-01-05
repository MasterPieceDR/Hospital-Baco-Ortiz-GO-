import React, { useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { AuthContext } from "../../context/AuthContext";

export default function Login() {
	const { login } = useContext(AuthContext);
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [role, setRole] = useState("");
	const [error, setError] = useState(null);
	const navigate = useNavigate();

	const onSubmit = async (e) => {
		e.preventDefault();
		setError(null);
		try {
			await login(username, password);
			navigate("/dashboard", { replace: true });
		} catch (err) {
			setError(err.response?.data?.error || err.message || "Login falló");
		}
	};

		return (
			<div className="container login-page">
				<div className="login-box">
					<h2>Iniciar sesión</h2>
					<div style={{ marginBottom: 18 }}>
						<label htmlFor="role-select">Rol:</label>
						<select
							id="role-select"
							style={{ width: "100%", padding: 8, borderRadius: 6, marginTop: 4, marginBottom: 8, border: "1px solid #e0e0e0" }}
							value={role}
							onChange={e => setRole(e.target.value)}
							required
						>
							<option value="" disabled>Seleccione un rol</option>
							<option value="Administrador">Administrador</option>
							<option value="Medico">Médico</option>
							<option value="Recepcion">Recepción</option>
							<option value="Paciente">Paciente</option>
							<option value="Enfermeria">Enfermería</option>
							<option value="Laboratorio">Laboratorio</option>
							<option value="Finanzas">Finanzas</option>
							<option value="Farmacia">Farmacia</option>
							<option value="IT_Admin">IT_Admin</option>
							<option value="Direccion">Dirección</option>
						</select>
					</div>
					<form onSubmit={onSubmit}>
						<label>Usuario</label>
						<input
							value={username}
							onChange={(e) => setUsername(e.target.value)}
							required
						/>
						<label>Contraseña</label>
						<input
							type="password"
							value={password}
							onChange={(e) => setPassword(e.target.value)}
							required
						/>
						{error && <div className="form-error">{error}</div>}
						<button className="login-submit" disabled={!role}>Entrar</button>
					</form>
				</div>
			</div>
		);
}
