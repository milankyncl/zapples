import React, {FC} from "react";
import classNames from "classnames";

export enum ButtonColor {
    Primary,
    Danger,
}

export enum ButtonSize {
    Xs,
    Sm,
    Md,
    Lg,
}

type ButtonType = 'submit' | 'reset' | 'button' | 'text';

interface Props {
    children: React.ReactNode;
    className?: string;
    onClick?: () => void;
    color?: ButtonColor,
    size?: ButtonSize;
    type?: ButtonType;
    disabled?: boolean;
}

function getColorClass (color: ButtonColor, type: ButtonType) {
    if (color === ButtonColor.Primary) {
        if (type === 'text') {
            return 'text-blue-500 enabled:hover:text-blue-600';
        }
        return 'bg-blue-500 enabled:hover:bg-blue-700 text-white';
    }
    if (color === ButtonColor.Danger) {
        if (type === 'text') {
            return 'text-red-500 enabled:hover:text-red-600';
        }
        return 'bg-red-500 enabled:hover:bg-red-600 text-white';
    }
}

function getSizeClass(size: ButtonSize, type: ButtonType) {
    if (size === ButtonSize.Xs) {
        if (type === 'text') {
            return 'text-xs font-medium py-1 px-1';
        }
        return 'text-sm font-medium py-1 px-2';
    }
    if (size === ButtonSize.Sm) {
        if (type === 'text') {
            return 'text-sm font-medium py-1 px-1';
        }
        return 'text-sm font-medium py-1 px-3';
    }
    if (size === ButtonSize.Md) {
        return 'text-md font-medium py-1 px-4';
    }
    if (size === ButtonSize.Lg) {
        return 'text-md font-medium py-2 px-5';
    }
}

export const Button: FC<Props> = ({
    children,
    className,
    color = ButtonColor.Primary,
    size = ButtonSize.Md,
    type = 'button',
    disabled,
    onClick,
}) => {
    return (
        <button
            onClick={onClick}
            type={type === 'text' ? 'button' : type}
            disabled={disabled}
            className={classNames(
                getColorClass(color, type),
                getSizeClass(size, type),
                "rounded transition-opacity disabled:opacity-60",
                className,
            )}
        >
            {children}
        </button>
    )
}