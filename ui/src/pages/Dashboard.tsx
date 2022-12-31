import React, {FC, useState} from 'react';
import useSWR from 'swr'
import {client} from '../api/client'
import {Button, ButtonSize} from "../components/atoms/Button";
import {PlusIcon} from "@heroicons/react/24/solid";
import {useNavigate} from "react-router-dom";
import {PageHeading} from "../components/atoms/PageHeading";
import {FeatureItem} from "./features/FeatureItem";
import {Feature} from "../api/models/feature";

const fetcher = (url: string) => client.get(url).then(res => res.data)

interface FeatureTogglesResponse {
    data: Feature[];
}

function noItemsYet(
    onAddClick: () => void,
) {
    return <>
        <p className="text-sm text-gray text-center italic">
            No features created yet
        </p>
        <div className="text-center">
            <Button onClick={onAddClick} className="inline-flex items-center">
                <PlusIcon className="mr-2 -ml-1 w-4 h-4 text-white" />
                <span>Add new feature</span>
            </Button>
        </div>
    </>;
}

export const Dashboard: FC = () => {
    const navigate = useNavigate();
    const { data, error, isLoading } = useSWR<FeatureTogglesResponse>('/api/features', fetcher)

    if (error) return <div>failed to load</div>
    if (isLoading || !data) return <div>loading...</div>

    const features = data.data;

    const onToggleCheck = (id: number) => async (enabled: boolean) => {
        await client.put(`/api/features/${id}/toggle`, {
            enabled,
        });
    }
    const createNewFeature = () => navigate('/features/create');

    return <>
        <div className="flex items-center mb-4">
            <PageHeading>Features</PageHeading>
            <div className="text-right flex-1">
                <Button size={ButtonSize.Sm} onClick={createNewFeature} className="inline-flex items-center">
                    <PlusIcon className="mr-2 -ml-1 w-4 h-4 text-white" />
                    <span>Create new</span>
                </Button>
            </div>
        </div>
        <div className="divide-y divide-gray-300/50">
            {features.length === 0 && noItemsYet(createNewFeature)}
            {features.map((feature) => <FeatureItem feature={feature} onToggle={onToggleCheck(feature.id)} />)}
        </div>
    </>
}