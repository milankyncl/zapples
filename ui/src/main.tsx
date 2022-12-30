import React from 'react'
import ReactDOM from 'react-dom/client'
import {
    createHashRouter,
    RouterProvider,
} from "react-router-dom";

import './index.css';
import {PageWrapper} from "./components/PageWrapper";
import {Dashboard} from "./pages/Dashboard";

const router = createHashRouter([
    {
        path: "/",
        element: <Dashboard />,
    },
]);

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <PageWrapper>
            <RouterProvider router={router} />
        </PageWrapper>
    </React.StrictMode>,
)
