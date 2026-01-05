import React from "react";

export default function NewsModal({ open, onClose, news }) {
  if (!open || !news) return null;
  return (
    <div style={{
      position: "fixed", inset: 0, background: "rgba(0,0,0,0.35)", zIndex: 1000, display: "flex", alignItems: "center", justifyContent: "center"
    }}>
      <div style={{
        background: "#fff", borderRadius: 16, maxWidth: 420, width: "90%", boxShadow: "0 8px 32px rgba(0,0,0,0.18)", padding: 28, position: "relative"
      }}>
        <button onClick={onClose} style={{ position: "absolute", top: 12, right: 16, fontSize: 22, background: "none", border: "none", cursor: "pointer", color: "#888" }}>&times;</button>
        <img src={news.img} alt={news.title} style={{ width: "100%", borderRadius: 10, marginBottom: 18, maxHeight: 180, objectFit: "cover" }} />
        <h2 style={{ fontSize: 22, fontWeight: 700, marginBottom: 8 }}>{news.title}</h2>
        <div style={{ color: "#17624a", fontWeight: 500, fontSize: 15, marginBottom: 8 }}>{news.date}</div>
        <div style={{ color: "#333", fontSize: 16, marginBottom: 10 }}>{news.text}</div>
        {news.body && <div style={{ color: "#444", fontSize: 15, marginTop: 10 }}>{news.body}</div>}
      </div>
    </div>
  );
}
