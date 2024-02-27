import React from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Home from '../Home/Home';
import { BlogPage } from '../Blog/BlogPage';

 function Router() {
  return (
    <>
    <BrowserRouter>

        <Routes>
            <Route path='/' element={<Home/>}/>
            <Route path='/blog' element={<BlogPage/>}/>

        </Routes>
    
    
    </BrowserRouter>
    
    
    </>
  )
}
export default Router;
