export interface Feature {
    id: number;
    key: string;
    description: string | null;
    enabled: boolean;
    enabledSince: Date | null;
    enabledUntil: Date | null;
    createdAt: Date;
}