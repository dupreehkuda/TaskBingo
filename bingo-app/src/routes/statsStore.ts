import { writable } from 'svelte/store';

export interface StatsWindow {
    days: number;
    from: string;
    to: string;
}

export interface StatsBucket {
    label: string;
    soloBingo: number;
    duoBingo: number;
    tasksClosed: number;
}

export interface StatsPackCount {
    id: string;
    title: string;
    count: number;
}

export interface StatsSummary {
    games: number;
    wins: number;
    losses: number;
    totalBingo: number;
    soloBingo: number;
    duoBingo: number;
    tasksClosed: number;
}

export interface StatsResult {
    window: StatsWindow;
    buckets: StatsBucket[];
    topPacks: StatsPackCount[];
    summary: StatsSummary;
}

const Stats = writable<StatsResult | null>(null);

export default Stats;
