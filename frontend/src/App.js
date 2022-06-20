import React from "react";
import { Routes, Route } from "react-router-dom";
import {
	Landing,
	Login,
	Register,
	Question,
	Language,
	HomePage,
} from "./Pages";

export const App = () => {
	return (
		<Routes>
			<Route path="/" element={<Landing />}>
				{" "}
			</Route>
			<Route path="/login" element={<Login />}></Route>
			<Route path="/register" element={<Register />}></Route>
			<Route path="/HomePage" element={<HomePage />}></Route>
		</Routes>
	);
};
export default App;
