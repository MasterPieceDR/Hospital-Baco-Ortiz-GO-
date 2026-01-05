import React, { createContext, useState, useEffect, useContext } from "react";
import api from "../api/axios";

export const AuthContext = createContext(null);

export function AuthProvider({ children }) {
	const [token, setToken] = useState(() => localStorage.getItem("token") || null);
	const [user, setUser] = useState(() => {
		try {
			const u = localStorage.getItem("user");
			return u ? JSON.parse(u) : null;
		} catch {
			return null;
		}
	});

	useEffect(() => {
		if (token) localStorage.setItem("token", token);
		else localStorage.removeItem("token");
	}, [token]);

	useEffect(() => {
		if (user) localStorage.setItem("user", JSON.stringify(user));
		else localStorage.removeItem("user");
	}, [user]);

	const login = async (username, password) => {
		const res = await api.post("/auth/login", { username, password });
		const data = res.data;
		if (!data || !data.token) throw new Error("Login falló");
		setToken(data.token);
		setUser({ username: data.username, roles: data.roles || [] });
		return data;
	};

	const logout = () => {
		setToken(null);
		setUser(null);
		localStorage.removeItem("token");
		localStorage.removeItem("user");
	};

	return (
		<AuthContext.Provider value={{ token, user, login, logout }}>
			{children}
		</AuthContext.Provider>
	);
}

export function useAuth() {
	return useContext(AuthContext);
}
