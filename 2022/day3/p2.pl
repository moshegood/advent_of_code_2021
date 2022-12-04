use strict;
use warnings;

sub valueOfItem{
    my $c = shift;
    return ord($c) - ord('a') + 1 if $c =~ m/[a-z]/;
    return ord($c) - ord('A') + 27 if $c =~ m/[A-Z]/;
    say("INVALID ITEM: $c");
    return 0;
}

my $sum = 0;
my @lines;
while (my $line = <>) {
    chomp $line;
    push (@lines,$line);
    next if (@lines != 3);

    my %counts;
    for my $i (0 .. 2) {
        my %seen;
        foreach my $c (split('', $lines[$i])) {
            $seen{$c} = 1;
        }
        foreach my $k (keys %seen){
            $counts{$k}++;
        }
    }
    foreach my $k (keys %counts) {
        $sum += valueOfItem($k) if $counts{$k} == 3;
    }
    @lines = ();
}

print "$sum\n";
