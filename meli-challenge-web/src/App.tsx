import './App.css';

import { Routes, Route } from "react-router-dom";
import HomePage from './pages/Home/Home';
import ItemPage from './pages/Item/Item';
import Items from './pages/Items/Items';
import NotFound from './pages/NotFound/NotFound';

import { Navbar } from './components';

function App() {
    return (
        <div className="App h-screen">
            <Navbar/>
            <Routes>
                <Route path="/" element={<HomePage />}/>
                <Route path="/items" element={<Items/>}/>
                <Route path="/item/:id" element={<ItemPage/>}/>
                <Route path="*" element={<NotFound/>} />
            </Routes>
        </div>
    );
}

export default App;

