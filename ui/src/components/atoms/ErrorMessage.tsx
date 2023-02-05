import React from 'react';

interface Props {
    children: React.ReactNode;
}

export function ErrorMessage({ children }: Props) {
    return <div className="py-4 flex items-center justify-center text-red-500 text-sm">
        <span className="mr-2">😓</span>
        <span>{children}</span>
    </div>
}