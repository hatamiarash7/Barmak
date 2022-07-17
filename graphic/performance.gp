# Graphic will be saved as 800x600 png image file
set terminal png

# Allows grid lines to be drawn on the plot
set grid

# Setting the graphic file name to be saved
set output graphic_file_name

# The graphic's main title
set title "Performance Comparison"

# Since the input file is a CSV file, we need to tell gnuplot that data fields are separated by comma
set datafile separator ","

# Disable key box
set key off

# Label for y axis
set ylabel y_label

# Range for values in y axis
set yrange[y_range_min:y_range_max]

# To avoid displaying large numbers in exponential format
set format y "%.0f"

# Vertical label for x values
set xtics rotate

# Set boxplots
set style fill solid
set boxwidth 0.5

# Plot graphic for each line of input file
plot for [i=0:*] file_path every ::i::i using column_1:column_2:xtic(2) with boxes
