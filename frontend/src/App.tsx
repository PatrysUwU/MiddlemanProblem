import { BrowserRouter, Route, Routes } from "react-router";
import SupplyDemandForm from "./Components/SupplyDemandForm";
import TransportForm from "./Components/TransportForm";
import ResultPage from "./Components/ResultPage";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<SupplyDemandForm />} />
        <Route path="/transportForm" element={<TransportForm />} />
        <Route path="/result" element={<ResultPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
