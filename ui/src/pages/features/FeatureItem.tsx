import React, {FC, useState} from "react";
import {Toggle} from "../../components/atoms/Toggle";
import {Feature} from "../../api/models/feature";
import {Button, ButtonColor, ButtonSize} from "../../components/atoms/Button";
import {PencilIcon, TrashIcon} from "@heroicons/react/20/solid";

interface Props {
    feature: Feature;
    onToggle: (toggled: boolean) => Promise<void>;
    onRemove: () => Promise<void>;
    onEdit: () => void;
}

export const FeatureItem: FC<Props> = ({ feature, onToggle, onRemove, onEdit }) => {
    const [loading, setLoading] = useState(false);

    const onToggleFeature = async (toggled: boolean) => {
        setLoading(true);
        try {
            await onToggle(toggled);
        } catch (e) {
            console.error(e);
        } finally {
            setLoading(false);
        }
    }
    const onRemoveFeature = async () => {
        setLoading(true);
        try {
            await onRemove();
        } catch (e) {
            console.error(e);
        } finally {
            setLoading(false);
        }
    }

    return <div className="py-2 pb-3 flex items-center">
        <div className="flex-1">
            <span className="font-medium">{feature.key}</span>
            {feature.description && <>
                <br />
                <span className="text-sm text-gray-500">{feature.description}</span>
            </>}
            <div className="flex mt-1">
                <Button
                    type="text"
                    className="inline-flex items-center mr-2"
                    size={ButtonSize.Xs}
                    color={ButtonColor.Primary}
                    disabled={loading}
                    onClick={onEdit}
                >
                    <PencilIcon className="mr-1 -ml-1 w-4 h-4 text-blue-500" />
                    Edit
                </Button>
                <Button
                    type="text"
                    className="inline-flex items-center"
                    size={ButtonSize.Xs}
                    color={ButtonColor.Danger}
                    disabled={loading}
                    onClick={onRemoveFeature}
                >
                    <TrashIcon className="mr-1 -ml-1 w-4 h-4 text-red-500" />
                    Remove
                </Button>
            </div>
        </div>
        <div>
            <Toggle defaultToggled={feature.enabled} onToggle={onToggleFeature} disabled={loading} />
        </div>
    </div>
}