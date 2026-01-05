// src/components/NewsSection.jsx
import news1 from "../assets/news1.jpg";
import news2 from "../assets/news2.jpg";
import news3 from "../assets/news3.jpg";
import React, { useState } from "react";
import NewsModal from "./NewsModal";

const news = [
	{
		img: news1,
		title: "Guerreros contra el cáncer son condecorados",
		date: "08 Noviembre 2022",
		text: "Reconocimiento a los pequeños pacientes que luchan día a día con valentía.",
	},
	{
		img: news2,
		title: "Nueva área de cuidados intensivos inaugurada",
		date: "15 Octubre 2022",
		text: "Ampliación de la atención crítica con equipos modernos.",
	},
	{
		img: news3,
		title: "Campaña de prevención de quemaduras en niños",
		date: "01 Octubre 2022",
		text: "La mayoría de accidentes son prevenibles; campaña busca crear conciencia.",
	},
];

export default function NewsSection() {
	const [modalOpen, setModalOpen] = useState(false);
	const [selectedNews, setSelectedNews] = useState(null);

	const openModal = (item) => {
		setSelectedNews(item);
		setModalOpen(true);
	};
	const closeModal = () => {
		setModalOpen(false);
		setSelectedNews(null);
	};

	return (
		<section className="news container py-12">
			<h2 className="section-title text-center text-2xl font-bold mb-8">
				Noticias
			</h2>
			<div className="news-grid grid grid-cols-1 md:grid-cols-3 gap-8">
				{news.map((item, idx) => (
					<div className="news-item rounded-xl shadow-md bg-white" key={item.title}>
						<img
							src={item.img}
							alt={item.title}
							className="w-full h-48 object-cover rounded-t-xl"
						/>
						<div className="p-4">
							<h3 className="font-semibold">{item.title}</h3>
							<p className="text-xs text-gray-500 mt-1">{item.date}</p>
							<p className="text-sm text-gray-600 mt-2">{item.text}</p>
							<button
								className="text-blue-600 text-sm mt-2 inline-block hover:underline"
								style={{ background: "none", border: "none", padding: 0, cursor: "pointer" }}
								onClick={() => openModal(item)}
							>
								Leer más &rarr;
							</button>
						</div>
					</div>
				))}
			</div>
			<NewsModal open={modalOpen} onClose={closeModal} news={selectedNews} />
		</section>
	);
}
