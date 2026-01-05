// src/pages/Home.jsx
import React from "react";
import HeroCarousel from "../components/HeroCarousel";
import ServicesSection from "../components/ServicesSection";
import NewsSection from "../components/NewsSection";

export default function Home() {
	return (
		<>
			<HeroCarousel />
			<ServicesSection />
			<NewsSection />
		</>
	);
}
