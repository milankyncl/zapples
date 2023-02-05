import React, {FC} from 'react';

interface Props {
    children?: React.ReactNode;
}

export const PageHeading: FC<Props> = ({ children }) => (
    <h1 className="font-medium leading-tight text-2xl mt-0 text-black-900">
        {children}
    </h1>
)