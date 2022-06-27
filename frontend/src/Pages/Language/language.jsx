import React from "react";
import "./language.css";
import API from "../../Api/Api";
import { useState, useEffect } from "react";
import axios from "axios";

import ListLanguage from "../../Components/Language/ListLanguage";

const Language = () => {
	const [listLanguage, setListLanguage] = useState([]);

	const fetchListPost = async () => {
		let auth = localStorage.getItem("token");

		// e.preventDefault();
		try {
			let res = await axios.get(`${API.API_URL}/home/start-ruang-arah`, {
				headers: {
					Accept: "/",
					"Content-Type": "application/json",
					Authorization: "Bearer " + auth,
				},
			});
			// console.log(res);
			setListLanguage(res.data.data);
		} catch (error) {
			console.log(error);
		}
	};

	useEffect(() => {
		fetchListPost();
	}, []);

	return (
		<section id="language-page">
			<div className="container">
				<h1>
					Bahasa <span> Pemrograman </span>{" "}
				</h1>
				<div className="card">
					<div className="card-content">
						{listLanguage.map((item) => (
							<ListLanguage
								key={item.id}
								image={item.url_images}
								language={item.name}
							/>
						))}
					</div>
				</div>
			</div>
		</section>
	);
};

export default Language;
