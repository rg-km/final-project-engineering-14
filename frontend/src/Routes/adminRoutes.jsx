import React, { useEffect, useContext } from "react";
import { useNavigate, Outlet } from "react-router-dom";
import AuthContext from "../store/auth-context";

function PrivateRoute() {
	const user = useContext(AuthContext);

	const navigate = useNavigate();

	useEffect(() => {
		const role = localStorage.getItem("role");

		if (!user.role === "admin" && !role) {
			navigate("/HomePage");
		}
	}, []);
	return (
		<div>
			<Outlet />
		</div>
	);
}

export default PrivateRoute;
