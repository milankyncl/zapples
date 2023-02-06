export interface Feature {
    id: number;
    key: string;
    description: string | null;
    enabled: boolean;
    enabledSince: string | null;
    enabledUntil: string | null;
    createdAt: string;
}