import { BrowserRouter, Route, Routes } from "react-router";
import SupplyDemandForm from "./Components/SupplyDemandForm";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<SupplyDemandForm />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
