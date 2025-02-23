import Layout from "./Components/Layout/Layout"
import Home from "./Components/Home/Home";
import { BrowserRouter, Routes, Route } from "react-router-dom"
import Messages from "./Components/Messages/Messages";
import Login from "./Components/Login/Login";
import Signup from "./Components/Signup/Signup";
import ProtectedRoute from "./Components/ProtectedRoute/ProtectedRoute";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login/>}/>
        <Route path="/signup" element={<Signup/>}/>
        <Route element={<ProtectedRoute />}>
          <Route path="/" element={<Layout />}>
            <Route index element={<Home />} />
            <Route path="/home" element={<Home />} />
            <Route path="/messages" element={<Messages />} />
          </Route>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App
